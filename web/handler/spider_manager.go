package handler

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-martini/martini"
	"github.com/huzorro/ospider/web/user" 
	"net/http"
	"log"
	"github.com/martini-contrib/sessions"
	"encoding/json"
	"strings"
	"strconv"
	"reflect"
	"net/url"
	"github.com/martini-contrib/render"
	"github.com/huzorro/ospider/util"
	"time"
	"github.com/huzorro/ospider/common"
)

type QueueName struct {
    Task string `json:"task"`
    Result string `json:"result"`
}
type Manager struct {
    Id int64 `json:"id"`
    Uid int64 `json:"uid"`
    Name string `json:"name"`
    QueueName
    QueueNameStr string `json:"queueName"`
    Status int64 `json:"status"`
    Logtime string `json:"logtime"`
}
type SpiderRelation struct {
    Manager
    user.SpStatUser 
}

func GetSpidersApi(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        manager Manager
        managers []Manager
		js   []byte
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
    sqlStr := `select id, uid, name, queue_name_json, status, logtime from spider_manager where status = 1`
    stmtOut, err := db.Prepare(sqlStr)                                                                                                                          

	rows, err := stmtOut.Query()
    defer func ()  {
        stmtOut.Close()
        rows.Close()
    }() 
	if err != nil {
		log.Printf("%s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	for rows.Next() {
        if err := rows.Scan(&manager.Id, &manager.Uid, &manager.Name, 
                            &manager.QueueNameStr, &manager.Status, &manager.Logtime); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
        if err := json.Unmarshal([]byte(manager.QueueNameStr), &manager.QueueName); err != nil {
            log.Printf("json Unmarshal fails %s", err)
            rs, _ := json.Marshal(s)
            return http.StatusOK, string(rs)            
        }
        managers = append(managers, manager)
	}
    
	if js, err = json.Marshal(managers); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)       
}

func GetSpider(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        manager Manager
		js   []byte
		spUser user.SpStatUser
		con  string        
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}

	switch spUser.Access.Rule {
	case user.GROUP_PRI_ALL:
	case user.GROUP_PRI_ALLOW:
		con = " a.uid IN(" + strings.Join(spUser.Access.Group, ",") + ") and "
	case user.GROUP_PRI_BAN:
		con = " a.uid NOT IN(" + strings.Join(spUser.Access.Group, ",") + ") and "
	default:
		log.Printf("group private erros")
	}
    sqlStr := `select id, uid, name, queue_name_json, status, logtime from spider_manager where ` + 
                con + " id = ?"
    stmtOut, err := db.Prepare(sqlStr)                                                                                                                          
	id, _ := strconv.Atoi(r.PostFormValue("Id"))

	rows, err := stmtOut.Query(id)
    defer func ()  {
        stmtOut.Close()
        rows.Close()
    }() 
	if err != nil {
		log.Printf("%s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	if rows.Next() {
        if err := rows.Scan(&manager.Id, &manager.Uid, &manager.Name, 
                            &manager.QueueNameStr, &manager.Status, &manager.Logtime); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
        if err := json.Unmarshal([]byte(manager.QueueNameStr), &manager.QueueName); err != nil {
            log.Printf("json Unmarshal fails %s", err)
            rs, _ := json.Marshal(s)
            return http.StatusOK, string(rs)            
        }
	}

	if js, err = json.Marshal(manager); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)
}

func EditSpider(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        manager Manager
		js   []byte
		spUser user.SpStatUser
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&manager).Elem()
	rValue := reflect.ValueOf(&manager).Elem()
	for i := 0; i < rType.NumField(); i++ {
		fN := rType.Field(i).Name
		if p, _ := url.QueryUnescape(strings.TrimSpace(r.PostFormValue(fN))); p != "" {
            switch rType.Field(i).Type.Kind() {
            case reflect.String:
                rValue.FieldByName(fN).SetString(p)
            case reflect.Int64:
                in, _ := strconv.ParseInt(p, 10, 64)
                rValue.FieldByName(fN).SetInt(in)
            case reflect.Float64:
                fl, _ := strconv.ParseFloat(p, 64)
                rValue.FieldByName(fN).SetFloat(fl)
            default:
                log.Printf("unknow type %s", p)
            }            
        }         
	}
    manager.Task, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Task"))) 
    manager.Result, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Result")))
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `update spider_manager set name = ?, queue_name_json = ?, status = ? 
                where id = ? and uid = ?` 
    stmtIn, err := db.Prepare(sqlStr)
    defer func ()  {
        stmtIn.Close()
    }()     
    if js, err = json.Marshal(manager.QueueName); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
    if _, err = stmtIn.Exec(manager.Name, js, manager.Status, manager.Id, spUser.Id); err != nil {
		log.Printf("update spider fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)         
    }
	s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js)           
}

func AddSpider(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string)  {
	var (
        manager Manager
		js   []byte
		spUser user.SpStatUser       
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&manager).Elem()
	rValue := reflect.ValueOf(&manager).Elem()
	for i := 0; i < rType.NumField(); i++ {
		fN := rType.Field(i).Name
		if p, _ := url.QueryUnescape(strings.TrimSpace(r.PostFormValue(fN))); p != "" {         
            switch rType.Field(i).Type.Kind() {
            case reflect.String:
                rValue.FieldByName(fN).SetString(p)
            case reflect.Int64:
                in, _ := strconv.ParseInt(p, 10, 64)
                rValue.FieldByName(fN).SetInt(in)
            case reflect.Float64:
                fl, _ := strconv.ParseFloat(p, 64)
                rValue.FieldByName(fN).SetFloat(fl)
            default:
                log.Printf("unknow type %s", p)
            }            
        }         
	}
    manager.Task, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Task"))) 
    manager.Result, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Result")))
    
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `insert into spider_manager(uid, name, queue_name_json, logtime) 
                values(?, ?, ?, ?)`
    stmtIn, err := db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        log.Printf("db prepare fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
    if js, err = json.Marshal(manager.QueueName); err != nil {
        log.Printf("json Marshal fails %s", err) 
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
    } 
    if _, err :=  stmtIn.Exec(spUser.Id, manager.Name, string(js), time.Now().Format("2006-01-02 15:04:05")); err != nil {
        log.Printf("stmtIn exec fails %s", err) 
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
    }
    s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js) 
}

func GetSpiders(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	 cfg *common.Cfg, session sessions.Session, ms []*user.SpStatMenu, render render.Render)  {
	var (
		spiderRelation  *SpiderRelation
		spiderRelations []*SpiderRelation
		menu          []*user.SpStatMenu
		spUser        user.SpStatUser
		con           string
		totalN        int64
		destPn        int64
	)
	path := r.URL.Path
	r.ParseForm()
	value := session.Get(user.SESSION_KEY_QUSER)

	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
		return
	}

	switch spUser.Access.Rule {
	case user.GROUP_PRI_ALL:
	case user.GROUP_PRI_ALLOW:
		con = "WHERE a.uid IN(" + strings.Join(spUser.Access.Group, ",") + ")"
	case user.GROUP_PRI_BAN:
		con = "WHERE a.uid NOT IN(" + strings.Join(spUser.Access.Group, ",") + ")"
	default:
		log.Printf("group private erros")
	}

	for _, elem := range ms {
		if (spUser.Role.Menu & elem.Id) == elem.Id {
			menu = append(menu, elem)
		}
	}
	stmtOut, err := db.Prepare("SELECT COUNT(*) FROM spider_manager a " + con)
	if err != nil {
		log.Printf("%s", err)
		http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
		return
	}
	row := stmtOut.QueryRow()
	if err = row.Scan(&totalN); err != nil {
		log.Printf("%s", err)
		http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
		return
	}
	//page
	if r.URL.Query().Get("p") != "" {
		destPn, _ = strconv.ParseInt(r.URL.Query().Get("p"), 10, 64)
	} else {
		destPn = 1
	}
    sqlStr := `select a.id, a.uid, b.username, a.name, a.queue_name_json, a.status, a.logtime  
                from spider_manager a left join sp_user b on a.uid = b.id 
                ` + con + `order by id desc limit ?, ?`  
    log.Printf("%s", sqlStr)
    stmtOut, err = db.Prepare(sqlStr)
    defer stmtOut.Close()
	rows, err := stmtOut.Query(cfg.PageSize*(destPn-1), cfg.PageSize)
    defer rows.Close()
	if err != nil {
		log.Printf("%s", err)
		http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
		return
	}
	for rows.Next() {
        spiderRelation = &SpiderRelation{}
		if err := rows.Scan(&spiderRelation.Manager.Id,&spiderRelation.Manager.Uid,
            &spiderRelation.UserName, &spiderRelation.Name, &spiderRelation.QueueNameStr,
            &spiderRelation.Manager.Status, &spiderRelation.Manager.Logtime); err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return
		}
        if err := json.Unmarshal([]byte(spiderRelation.QueueNameStr), &spiderRelation.QueueName); err != nil {
            log.Printf("%s", err)
            http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return            
        }
		spiderRelations = append(spiderRelations, spiderRelation)
	}
    
	paginator := util.NewPaginator(r, cfg.PageSize, totalN)

	ret := struct {
		Menu      []*user.SpStatMenu
		Result    []*SpiderRelation
		Paginator *util.Paginator
		User      *user.SpStatUser
	}{menu, spiderRelations, paginator, &spUser}
	if path == "/" {
		render.HTML(200, "spiderview", ret)
    } else {
        index := strings.LastIndex(path, "/")
        render.HTML(200, path[index+1:], ret)       
    }
	
}