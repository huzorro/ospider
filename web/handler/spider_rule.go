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

type Selector struct {
    Title string `json:"title"`
    Content string `json:"content"`
    Section string `json:"section"`
    Href string `json:"href"`
    Filter string `json:"filter"`
}

type Rule struct {
    Id int64 `json:"id"`
    Uid int64 `json:"uid"`
    Name string `json:"name"`
    SpiderId int64 `json:"spiderid"`
    Manager Manager `json:"manager"`
    Url string `json:"url"`
    Selector Selector `json:"selector"`
    SelectorStr string `json:"selectorStr"`
    Status int64 `json:"status"`
    Logtime string `json:"logtime"`
}
type RuleRelation struct {
    Rule
    user.SpStatUser 
}
func GetRulesApi(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
		spUser        user.SpStatUser        
        rule Rule
        rules []Rule
		js   []byte     
        con  string
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	value := session.Get(user.SESSION_KEY_QUSER)

	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
        rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}

	switch spUser.Access.Rule {
	case user.GROUP_PRI_ALL:
	case user.GROUP_PRI_ALLOW:
		con = "WHERE uid IN(" + strings.Join(spUser.Access.Group, ",") + ") and status = 1"
	case user.GROUP_PRI_BAN:
		con = "WHERE uid NOT IN(" + strings.Join(spUser.Access.Group, ",") + ") and status =1"
	default:
		log.Printf("group private erros")
	}
    sqlStr := `select id, uid, name, spiderid, url, rule_json,status, logtime from spider_rule ` + con
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
        if err := rows.Scan(&rule.Id, &rule.Uid, &rule.Name, 
                            &rule.SpiderId, &rule.Url, &rule.SelectorStr, &rule.Status, &rule.Logtime); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
        if err := json.Unmarshal([]byte(rule.SelectorStr), &rule.Selector); err != nil {
            log.Printf("json Unmarshal fails %s", err)
            rs, _ := json.Marshal(s)
            return http.StatusOK, string(rs)
        }
        rules = append(rules, rule)
	}
    
	if js, err = json.Marshal(rules); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)        
}

func GetRule(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        rule Rule
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
    sqlStr := `select a.id, a.uid, a.name, a.spiderid, c.name, a.url, a.rule_json, a.status, a.logtime 
               from spider_rule a left join spider_manager c on a.spiderid = c.id where ` + 
               con + " a.id = ?"
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
        if err := rows.Scan(&rule.Id, &rule.Uid, &rule.Name,&rule.SpiderId, &rule.Manager.Name, &rule.Url, 
                            &rule.SelectorStr, &rule.Status, &rule.Logtime); err != nil {
           log.Printf("%s", err)
           rs, _:= json.Marshal(s)
           return http.StatusOK, string(rs)
        }
        if err := json.Unmarshal([]byte(rule.SelectorStr), &rule.Selector); err != nil {
            log.Printf("json Unmarshal fails %s", err)
            rs, _ := json.Marshal(s)
            return http.StatusOK, string(rs)            
        }
	}

	if js, err = json.Marshal(rule); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	return http.StatusOK, string(js)
}

func EditRule(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string) {
	var (
        rule Rule
		js   []byte
		spUser user.SpStatUser
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&rule).Elem()
	rValue := reflect.ValueOf(&rule).Elem()
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
    rule.Selector.Title, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Title")))
    rule.Selector.Content, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Content")))
    rule.Selector.Section, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Section")))
    rule.Selector.Href, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Href")))
    rule.Selector.Filter, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Filter")))
    if len(rule.Selector.Title) <= 0 || len(rule.Selector.Content) <= 0 || len(rule.Selector.Section) <= 0 ||
        len(rule.Name) <= 0 || len(rule.Url) <= 0 {
        log.Printf("post From value [%s] is empty", "Title//Content//Section//Name//Url//Charset")
        rs, _ := json.Marshal(s)
        return http.StatusOK, string(rs)             
    }
           
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `update spider_rule set name = ?, spiderid = ?, url = ?, 
               rule_json = ?, status = ? where id = ? and uid = ?`       
    stmtIn, err := db.Prepare(sqlStr)
    defer func ()  {
        stmtIn.Close()
    }()     
    if js, err = json.Marshal(rule.Selector); err != nil {
		log.Printf("json Marshal fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)         
    } 
    if _, err = stmtIn.Exec(rule.Name, rule.SpiderId, rule.Url, js, 
                            rule.Status, rule.Id, spUser.Id); err != nil {
		log.Printf("update spider rule fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)         
    }
	s.Status = "200"
    s.Text = "操作成功"
    js, _ = json.Marshal(s)
    return http.StatusOK, string(js)           
}

func AddRule(r *http.Request, w http.ResponseWriter, db *sql.DB,
	log *log.Logger, session sessions.Session, p martini.Params) (int, string)  {
	var (
        rule Rule
		js   []byte
		spUser user.SpStatUser       
	)
    s := user.Status{Status:"500", Text:"操作失败"}
	r.ParseForm()
	rType := reflect.TypeOf(&rule).Elem()
	rValue := reflect.ValueOf(&rule).Elem()
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

    rule.Selector.Title, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Title")))
    rule.Selector.Content, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Content")))
    rule.Selector.Section, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Section")))
    rule.Selector.Href, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Href")))
    rule.Selector.Filter, _ = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("Filter")))    
    if len(rule.Selector.Title) <= 0 || len(rule.Selector.Content) <= 0 || len(rule.Selector.Section) <= 0 ||
        len(rule.Name) <= 0 || len(rule.Url) <= 0 {
        log.Printf("post From value [%s] is empty", "Title//Content//Section//Name//Url//Charset")
        rs, _ := json.Marshal(s)
        return http.StatusOK, string(rs)             
    }        
	value := session.Get(user.SESSION_KEY_QUSER)
	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &spUser)
	} else {
		log.Printf("session stroe type error")
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
    sqlStr := `insert into spider_rule (uid, name, spiderid, url, rule_json, logtime) 
                values(?, ?, ?, ?, ?, ?)`
    stmtIn, err := db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        log.Printf("db prepare fails %s", err)
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
    if js, err = json.Marshal(rule.Selector); err != nil {
        log.Printf("json Marshal fails %s", err) 
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)        
    }
    if _, err :=  stmtIn.Exec(spUser.Id, rule.Name, rule.SpiderId, rule.Url, js, 
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

func GetRules(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	 cfg *common.Cfg, session sessions.Session, ms []*user.SpStatMenu, render render.Render)  {
	var (
		ruleRelation  *RuleRelation
		ruleRelations []*RuleRelation
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
	stmtOut, err := db.Prepare("SELECT COUNT(*) FROM spider_rule a " + con)
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
    sqlStr := `select a.id, a.uid, b.username, a.name, a.spiderid, c.name, a.url, a.rule_json, a.status, a.logtime  
                from spider_rule a left join sp_user b on a.uid = b.id 
                left join spider_manager c on a.spiderid = c.id 
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
        ruleRelation = &RuleRelation{}
		if err := rows.Scan(&ruleRelation.Rule.Id,&ruleRelation.Rule.Uid,
            &ruleRelation.SpStatUser.UserName, &ruleRelation.Rule.Name, 
            &ruleRelation.Rule.SpiderId, &ruleRelation.Rule.Manager.Name, 
            &ruleRelation.Rule.Url, &ruleRelation.Rule.SelectorStr, 
            &ruleRelation.Rule.Status, &ruleRelation.Rule.Logtime); err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return
		}
        if err := json.Unmarshal([]byte(ruleRelation.Rule.SelectorStr), &ruleRelation.Rule.Selector); err != nil {
            log.Printf("%s", err)
            http.Redirect(w, r, user.ERROR_PAGE_NAME, 301)
			return            
        }
		ruleRelations = append(ruleRelations, ruleRelation)
	}
	paginator := util.NewPaginator(r, cfg.PageSize, totalN)

	ret := struct {
		Menu      []*user.SpStatMenu
		Result    []*RuleRelation
		Paginator *util.Paginator
		User      *user.SpStatUser
	}{menu, ruleRelations, paginator, &spUser}
	if path == "/" {
		render.HTML(200, "ruleview", ret)
    } else {
        index := strings.LastIndex(path, "/")
        render.HTML(200, path[index+1:], ret)       
    }
	
}