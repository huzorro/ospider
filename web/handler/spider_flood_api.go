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
    "encoding/binary"

)

type FloodApi struct {
    Id int64 `json:"id"`
    Uid int64 `json:"uid"`
    Name string `json:"name"`
    Api string `json:"api"`
    Powerlevel int64 `json:"powerlevel"`
    Time int64 `json:"time"`    
    Status int64 `json:"status"`
    Uptime int64 `json:"uptime"`
    Logtime string `json:"logtime"`
}
type FloodApiRelation struct {
    FloodApi
    user.SpStatUser 
}

func GetFloodApis(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        floodApi FloodApi
        floodApis []FloodApi
		js   []byte
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
    sqlStr := `select id, uid, name, api, powerlevel, time, status, uptime, logtime from spider_flood_api where status = 1`
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
        if err := rows.Scan(&floodApi.Id, &floodApi.Uid, &floodApi.Name, 
                            &floodApi.Api, &floodApi.Powerlevel, &floodApi.Time, &floodApi.Status, &floodApi.Uptime, &floodApi.Time); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
        floodApis = append(floodApis, floodApi)
	}
    
	if js, err = json.Marshal(floodApis); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)       
}

func GetFloodApi(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        floodApi FloodApi
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
    sqlStr := `select id, uid, name, api, powerlevel, time, status, uptime, logtime from spider_flood_api where ` + 
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
        if err := rows.Scan(&floodApi.Id, &floodApi.Uid, &floodApi.Name, 
                            &floodApi.Api, &floodApi.Powerlevel, &floodApi.Time, &floodApi.Status,&floodApi.Uptime,  &floodApi.Logtime); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
	}

	if js, err = json.Marshal(floodApi); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)
}

func EditFloodApi(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        floodApi FloodApi
		js   []byte
		spUser user.SpStatUser
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&floodApi).Elem()
	rValue := reflect.ValueOf(&floodApi).Elem()
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
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    

    sqlStr := `update spider_flood_api set name = ?, api = ?, powerlevel = ?, status = ?, time = ?   
                where id = ? and uid = ?` 
    stmtIn, err := db.Prepare(sqlStr)
    defer func ()  {
        stmtIn.Close()
    }()     

    if _, err = stmtIn.Exec(floodApi.Name, floodApi.Api, floodApi.Powerlevel, floodApi.Status, floodApi.Time, floodApi.Id, spUser.Id); err != nil {
		log.Printf("update spider fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)         
    }
	s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js)           
}

func AddFloodApi(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string)  {
	var (
        floodApi FloodApi
		js   []byte
		spUser user.SpStatUser       
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&floodApi).Elem()
	rValue := reflect.ValueOf(&floodApi).Elem()
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
    
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `insert into spider_flood_api(uid, name, api, powerlevel, time) 
                values(?, ?, ?, ?, ?)`
    stmtIn, err := db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        log.Printf("db prepare fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
	var powerlevelBuf = make([]byte, 8)
	binary.BigEndian.PutUint32(powerlevelBuf[:4], uint32(floodApi.Powerlevel))
	binary.BigEndian.PutUint32(powerlevelBuf[4:], uint32(floodApi.Powerlevel))
    newPowerlevel := binary.BigEndian.Uint64(powerlevelBuf)
    var timeBuf = make([]byte, 8)
    
    binary.BigEndian.PutUint32(timeBuf[:4], uint32(floodApi.Time))
    binary.BigEndian.PutUint32(timeBuf[4:], uint32(floodApi.Time))
    newTime := binary.BigEndian.Uint64(timeBuf)
    if _, err :=  stmtIn.Exec(spUser.Id, floodApi.Name, floodApi.Api, newPowerlevel, newTime); err != nil {
        log.Printf("stmtIn exec fails %s", err) 
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
    }
    s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js) 
}

func GetFloods(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	 cfg *common.Cfg, session sessions.Session, ms []*user.SpStatMenu, render render.Render)  {
	var (
		floodApiRelation  *FloodApiRelation
		floodApiRelations []*FloodApiRelation
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
	stmtOut, err := db.Prepare("SELECT COUNT(*) FROM spider_flood_api a " + con)
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
    sqlStr := `select a.id, a.uid, b.username, a.name, a.api, a.powerlevel, a.time, a.status, a.logtime  
                from spider_flood_api a left join sp_user b on a.uid = b.id 
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
        floodApiRelation = &FloodApiRelation{}
		if err := rows.Scan(&floodApiRelation.FloodApi.Id,&floodApiRelation.FloodApi.Uid,
            &floodApiRelation.UserName, &floodApiRelation.FloodApi.Name, &floodApiRelation.FloodApi.Api,
            &floodApiRelation.FloodApi.Powerlevel, &floodApiRelation.FloodApi.Time,
            &floodApiRelation.FloodApi.Status, &floodApiRelation.FloodApi.Logtime); err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return
		}
		floodApiRelations = append(floodApiRelations, floodApiRelation)
	}
    
	paginator := util.NewPaginator(r, cfg.PageSize, totalN)

	ret := struct {
		Menu      []*user.SpStatMenu
		Result    []*FloodApiRelation
		Paginator *util.Paginator
		User      *user.SpStatUser
	}{menu, floodApiRelations, paginator, &spUser}
	if path == "/" {
		render.HTML(200, "spiderview", ret)
    } else {
        index := strings.LastIndex(path, "/")
        render.HTML(200, path[index+1:], ret)       
    }
	
}