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

type DocumentSet struct {
    Position int64 `json:"position"`
    Display int64 `json:"display"`
    Member int64 `json:"uid"`
    NickName string `json:"nickname"`
    Check int64 `json:"check"` 
    Category int64 `json:"category_id"`
    CategoryTitle string `json:"title"`
    GroupId int64 `json:"group_id"`
    ModelId int64 `json:"model_id"` 
}

type Site struct {
    Id int64 `json:"id"`
    Uid int64 `json:"uid"`
    Name string `json:"name"`
    Url string `json:"url"`
    DocumentSet DocumentSet `json:"document_set"`
    DocumentSetStr string `json:"documentSetStr"`
    Rule Rule `json:"Rule"`
    Crontab Crontab `json:"Cron"`
    Status int64 `json:"status"`
    Logtime string `json:"logtime"`
}
type SiteRelation struct {
    Site
    user.SpStatUser 
}


func GetSite(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        site Site
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
		con = " a.uid IN(" + strings.Join(spUser.Access.Group, ",") + ")  and "
	case user.GROUP_PRI_BAN:
		con = " a.uid NOT IN(" + strings.Join(spUser.Access.Group, ",") + ")  and "
	default: 
		log.Printf("group private erros")
	}
    // sqlStr := `select a.id, a.uid, a.name, a.spiderid, c.name, a.url, a.rule_json, a.status, a.logtime 
    //            from spider_rule a left join spider_manager c on a.spiderid = c.id where ` + 
    //            con + " a.id = ?"
    sqlStr := `select a.id, a.uid, a.name, a.url, a.document_set, a.ruleid, b.name, a.cronid, c.name,
                a.status, a.logtime from spider_site a left join spider_rule b on a.ruleid = b.id 
                left join spider_crontab c on a.cronid = c.id where ` + con + " a.id = ?"
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
        if err := rows.Scan(&site.Id, &site.Uid, &site.Name, &site.Url, &site.DocumentSetStr, &site.Rule.Id, &site.Rule.Name, 
                            &site.Crontab.Id, &site.Crontab.Name, &site.Status, &site.Logtime); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
        if err := json.Unmarshal([]byte(site.DocumentSetStr), &site.DocumentSet); err != nil {
            log.Printf("json Unmarshal fails %s", err)
            rs, _ := json.Marshal(s)
            return http.StatusOK, string(rs)            
        }
	}

	if js, err = json.Marshal(site); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)
}

func EditSite(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        site Site
		js   []byte
		spUser user.SpStatUser
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&site).Elem()
	rValue := reflect.ValueOf(&site).Elem()
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
    site.DocumentSet.Position, _ = strconv.ParseInt(r.PostFormValue("Position"), 10, 64)
    site.DocumentSet.Display, _ = strconv.ParseInt(r.PostFormValue("Display"), 10, 64)
    site.DocumentSet.Member, _ = strconv.ParseInt(r.PostFormValue("Member"),10, 64)
    site.DocumentSet.NickName, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("NickName")))
    site.DocumentSet.Check, _ = strconv.ParseInt(r.PostFormValue("Check"), 10, 64)
    site.DocumentSet.Category, _ = strconv.ParseInt(r.PostFormValue("Category"), 10, 64)
    site.DocumentSet.CategoryTitle, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("CategoryTitle")))
    site.DocumentSet.GroupId, _ = strconv.ParseInt(r.PostFormValue("GroupId"), 10, 64)
    site.DocumentSet.ModelId, _ = strconv.ParseInt(r.PostFormValue("ModelId"), 10, 64)
    site.Rule.Id, _  = strconv.ParseInt(r.PostFormValue("RuleId"), 10, 64)
    site.Crontab.Id, _ = strconv.ParseInt(r.PostFormValue("CronId"), 10, 64)
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `update spider_site set name = ?, url = ?, document_set = ?, ruleid = ?,  
               cronid = ?, status = ? where id = ? and uid = ?`       
    stmtIn, err := db.Prepare(sqlStr)
    defer func ()  {
        stmtIn.Close()
    }()     
    if js, err = json.Marshal(site.DocumentSet); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)         
    } 
    if _, err = stmtIn.Exec(site.Name, site.Url, js, site.Rule.Id, site.Crontab.Id, 
                            site.Status, site.Id, spUser.Id); err != nil {
		log.Printf("update crontab fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)         
    }
	s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js)           
}

func AddSite(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string)  {
	var (
        site Site
		js   []byte
		spUser user.SpStatUser       
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&site).Elem()
	rValue := reflect.ValueOf(&site).Elem()
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

    site.DocumentSet.Position, _ = strconv.ParseInt(r.PostFormValue("Position"), 10, 64)
    site.DocumentSet.Display, _ = strconv.ParseInt(r.PostFormValue("Display"), 10, 64)
    site.DocumentSet.Member, _ = strconv.ParseInt(r.PostFormValue("Member"),10, 64)
    site.DocumentSet.NickName, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("NickName")))    
    site.DocumentSet.Check, _ = strconv.ParseInt(r.PostFormValue("Check"), 10, 64)
    site.DocumentSet.Category, _ = strconv.ParseInt(r.PostFormValue("Category"), 10, 64)
    site.DocumentSet.CategoryTitle, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("CategoryTitle")))
    site.DocumentSet.GroupId, _ = strconv.ParseInt(r.PostFormValue("GroupId"), 10, 64)
    site.DocumentSet.ModelId, _ = strconv.ParseInt(r.PostFormValue("ModelId"), 10, 64)
    site.Rule.Id, _  = strconv.ParseInt(r.PostFormValue("RuleId"), 10, 64)
    site.Crontab.Id, _ = strconv.ParseInt(r.PostFormValue("CronId"), 10, 64)        
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `insert into spider_site (uid, name, url, document_set, ruleid, cronid, logtime) 
                values(?, ?, ?, ?, ?, ?, ?)`
    stmtIn, err := db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        log.Printf("db prepare fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
    if js, err = json.Marshal(site.DocumentSet); err != nil {
        log.Printf("json Marshal fails %s", err) 
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
    if _, err :=  stmtIn.Exec(spUser.Id, site.Name, site.Url, js, site.Rule.Id, site.Crontab.Id, 
                                time.Now().Format("2006-01-02 15:04:05")); err != nil {
        log.Printf("stmtIn exec fails %s", err) 
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
    }
    s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js) 
}

func GetSites(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	 cfg *common.Cfg, session sessions.Session, ms []*user.SpStatMenu, render render.Render)  {
	var (
		siteRelation  *SiteRelation
		siteRelations []*SiteRelation
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
	stmtOut, err := db.Prepare("SELECT COUNT(*) FROM spider_site a " + con)
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
    sqlStr := `select a.id, a.uid, d.username, a.name, a.url, a.document_set, a.ruleid, b.name, a.cronid, c.name,
                a.status, a.logtime from spider_site a left join spider_rule b on a.ruleid = b.id 
                left join spider_crontab c on a.cronid = c.id 
                left join sp_user d on a.uid = d.id ` + con + " order by a.id desc limit ?, ?"    
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
        siteRelation = &SiteRelation{}
        if err := rows.Scan(&siteRelation.Site.Id, &siteRelation.Site.Uid,
                            &siteRelation.SpStatUser.UserName, &siteRelation.Site.Name,
                            &siteRelation.Site.Url, &siteRelation.DocumentSetStr,
                            &siteRelation.Rule.Id, &siteRelation.Rule.Name,
                            &siteRelation.Crontab.Id, &siteRelation.Crontab.Name,
                            &siteRelation.Status, &siteRelation.Logtime); err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return                                     
        }
        if err := json.Unmarshal([]byte(siteRelation.DocumentSetStr), &siteRelation.DocumentSet); err != nil {
            log.Printf("%s", err)
            http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return            
        }
		siteRelations = append(siteRelations, siteRelation)
	}
	paginator := util.NewPaginator(r, cfg.PageSize, totalN)

	ret := struct {
		Menu      []*user.SpStatMenu
		Result    []*SiteRelation
		Paginator *util.Paginator
		User      *user.SpStatUser
	}{menu, siteRelations, paginator, &spUser}
	if path == "/" {
		render.HTML(200, "siteview", ret)
    } else {
        index := strings.LastIndex(path, "/")
        render.HTML(200, path[index+1:], ret)       
    }
	
}