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
	"github.com/huzorro/ospider/common"
	"time"
)

type FloodTarget struct {
    Id int64 `json:"id"`
    Uid int64 `json:"uid"`       
    Url string `json:"url"`
    Host string `json:"host"`
    Port string `json:"port"`
    Method string `json:"method"`    
    Powerlevel int64 `json:"powerlevel"`
    Time int64 `json:"time"` 
    Crontab Crontab `json:"Cron"`
    Status int64 `json:"status"`
    Uptime string `json:"uptime"`
    Logtime string `json:"logtime"`
}
type FloodTargetRelation struct {
    FloodTarget
    user.SpStatUser 
}


func GetFloodTarget(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        floodTarget FloodTarget
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
    sqlStr := `select a.id, a.uid, a.url, a.host, a.port, a.method, a.powerlevel, a.time, a.cronid,  
                        b.name, a.status, a.uptime, a.logtime from spider_flood_target a 
                        left join spider_crontab b on a.cronid = b.id where ` + con + " a.id = ?"
                        
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
        if err := rows.Scan(&floodTarget.Id, &floodTarget.Uid, &floodTarget.Url, 
                            &floodTarget.Host, &floodTarget.Port, &floodTarget.Method, 
                            &floodTarget.Powerlevel, &floodTarget.Time, &floodTarget.Crontab.Id, &floodTarget.Crontab.Name, 
                            &floodTarget.Status,&floodTarget.Uptime,  &floodTarget.Logtime); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
	}

	if js, err = json.Marshal(floodTarget); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)
}

func EditfloodTarget(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        floodTarget FloodTarget
		js   []byte
		spUser user.SpStatUser
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&floodTarget).Elem()
	rValue := reflect.ValueOf(&floodTarget).Elem()
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
    floodTarget.Crontab.Id, _ = strconv.ParseInt(r.PostFormValue("CronId"), 10, 64)
    
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    

    sqlStr := `update spider_flood_target set url = ?, host = ?, cronid = ?, status = ?    
                where id = ? and uid = ?` 
    stmtIn, err := db.Prepare(sqlStr)
    defer func ()  {
        stmtIn.Close()
    }()     

    if _, err = stmtIn.Exec(floodTarget.Url, floodTarget.Host, floodTarget.Crontab.Id, 
                            floodTarget.Status, floodTarget.Id, spUser.Id); err != nil {
		log.Printf("update spider fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)         
    }
	s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js)           
}

func AddfloodTarget(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string)  {
	var (
        floodTarget FloodTarget
		js   []byte
		spUser user.SpStatUser       
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&floodTarget).Elem()
	rValue := reflect.ValueOf(&floodTarget).Elem()
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
    floodTarget.Crontab.Id, _ = strconv.ParseInt(r.PostFormValue("CronId"), 10, 64)    
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `insert into spider_flood_target(uid, url, host, cronid, logtime)
                values(?, ?, ?, ?, ?)`
    stmtIn, err := db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        log.Printf("db prepare fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
    if _, err :=  stmtIn.Exec(spUser.Id, floodTarget.Url, floodTarget.Host, 
                                floodTarget.Crontab.Id, time.Now().Format("2006-01-02 15:04:05")); err != nil {
        log.Printf("stmtIn exec fails %s", err) 
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
    }
    s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js) 
}

func GetfloodTargets(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	 cfg *common.Cfg, session sessions.Session, ms []*user.SpStatMenu, render render.Render)  {
	var (
		floodTargetRelation  *FloodTargetRelation
		floodTargetRelations []*FloodTargetRelation
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
	stmtOut, err := db.Prepare("SELECT COUNT(*) FROM spider_flood_target a " + con)
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
    sqlStr := `select a.id, a.uid, b.username, a.url, a.host, a.cronid, c.name, a.status, a.logtime  
                from spider_flood_target a left join sp_user b on a.uid = b.id 
                left join spider_crontab c on a.cronid = c.id  
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
        floodTargetRelation = &FloodTargetRelation{}
		if err := rows.Scan(&floodTargetRelation.FloodTarget.Id,&floodTargetRelation.FloodTarget.Uid,
            &floodTargetRelation.SpStatUser.UserName, &floodTargetRelation.FloodTarget.Url, 
            &floodTargetRelation.FloodTarget.Host, &floodTargetRelation.FloodTarget.Crontab.Id,
            &floodTargetRelation.FloodTarget.Crontab.Name, &floodTargetRelation.FloodTarget.Status,
            &floodTargetRelation.FloodTarget.Logtime); err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return
		}
		floodTargetRelations = append(floodTargetRelations, floodTargetRelation)
	}
    
	paginator := util.NewPaginator(r, cfg.PageSize, totalN)

	ret := struct {
		Menu      []*user.SpStatMenu
		Result    []*FloodTargetRelation
		Paginator *util.Paginator
		User      *user.SpStatUser
	}{menu, floodTargetRelations, paginator, &spUser}
	if path == "/" {
		render.HTML(200, "floodtargetview", ret)
    } else {
        index := strings.LastIndex(path, "/")
        render.HTML(200, path[index+1:], ret)       
    }
	
}