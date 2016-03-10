package user

import (
	"database/sql"
	"encoding/json"
	"github.com/huzorro/spfactor/sexredis"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"github.com/huzorro/ospider/util"
	"github.com/huzorro/ospider/common"
)



type Status struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

// type Cfg struct {
// 	//数据库类型
// 	Dbtype string `json:"dbtype"`
// 	//数据库连接uri
// 	Dburi string `json:"dburi"`
// 	//页宽
// 	PageSize int64 `json:"pageSize"`    
// }

type UserRelation struct {
	User   SpStatUser
	Role   SpStatRole
	Access SpStatAccess
}


func Logout(r *http.Request, w http.ResponseWriter, log *log.Logger, session sessions.Session) {
	session.Clear()
	http.Redirect(w, r, LOGIN_PAGE_NAME, 301)
}
func LoginCheck(r *http.Request, w http.ResponseWriter, log *log.Logger, db *sql.DB, session sessions.Session) (int, string) {
	//cross domain
	w.Header().Set("Access-Control-Allow-Origin", "*")
	un := r.PostFormValue("username")
	pd := r.PostFormValue("password")
	var (
		s Status
	)

	stmtOut, err := db.Prepare(`SELECT a.id, a.username, a.password, a.roleid, b.name, b.privilege, b.menu, 
		a.accessid, c.pri_group, c.pri_rule FROM sp_user a 
		INNER JOIN sp_role b ON a.roleid = b.id 
		INNER JOIN sp_access_privilege c ON a.accessid = c.id 
		WHERE username = ? AND password = ? `)

	if err != nil {
		log.Printf("get login user fails %s", err)
		s = Status{"500", "内部错误导致登录失败."}
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	result, err := stmtOut.Query(un, pd)
	defer func() {
		stmtOut.Close()
		result.Close()
	}()
	if err != nil {
		log.Printf("%s", err)
		//		http.Redirect(w, r, ERROR_PAGE_NAME, 301)
		s = Status{"500", "内部错误导致登录失败."}
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}
	if result.Next() {
		u := SpStatUser{}
		u.Role = &SpStatRole{}
		u.Access = &SpStatAccess{}
		var g string
		if err := result.Scan(&u.Id, &u.UserName, &u.Password, &u.Role.Id, &u.Role.Name, &u.Role.Privilege, &u.Role.Menu, &u.Access.Id, &g, &u.Access.Rule); err != nil {
			log.Printf("%s", err)
			s = Status{"500", "内部错误导致登录失败."}
			rs, _ := json.Marshal(s)
			return http.StatusOK, string(rs)
		} else {
			u.Access.Group = strings.Split(g, ";")
			//
			uSession, _ := json.Marshal(u)
			session.Set(SESSION_KEY_QUSER, uSession)
			s = Status{"200", "登录成功"}
			rs, _ := json.Marshal(s)
			return http.StatusOK, string(rs)
		}

	} else {
		log.Printf("%s", err)
		s = Status{"403", "登录失败,用户名/密码错误"}
		rs, _ := json.Marshal(s)
		return http.StatusOK, string(rs)
	}

}



func ViewUsersAction(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	redisPool *sexredis.RedisPool, cfg *common.Cfg, session sessions.Session, ms []*SpStatMenu, render render.Render) {
	var (
		userRelation  *UserRelation
		userRelations []*UserRelation
		menu          []*SpStatMenu
		user          SpStatUser
		con           string
		totalN        int64
		destPn        int64
	)
	path := r.URL.Path
	r.ParseForm()
	value := session.Get(SESSION_KEY_QUSER)

	if v, ok := value.([]byte); ok {
		json.Unmarshal(v, &user)
	} else {
		log.Printf("session stroe type error")
		http.Redirect(w, r, ERROR_PAGE_NAME, 301)
		return
	}

	switch user.Access.Rule {
	case GROUP_PRI_ALL:
	case GROUP_PRI_ALLOW:
		con = "WHERE a.id IN(" + strings.Join(user.Access.Group, ",") + ")"
	case GROUP_PRI_BAN:
		con = "WHERE a.id NOT IN(" + strings.Join(user.Access.Group, ",") + ")"
	default:
		log.Printf("group private erros")
	}

	for _, elem := range ms {
		if (user.Role.Menu & elem.Id) == elem.Id {
			menu = append(menu, elem)
		}
	}
	stmtOut, err := db.Prepare("SELECT COUNT(*) FROM sp_user a " + con)
	if err != nil {
		log.Printf("%s", err)
		http.Redirect(w, r, ERROR_PAGE_NAME, 301)
		return
	}
	row := stmtOut.QueryRow()
	if err = row.Scan(&totalN); err != nil {
		log.Printf("%s", err)
		http.Redirect(w, r, ERROR_PAGE_NAME, 301)
		return
	}
	//page
	if r.URL.Query().Get("p") != "" {
		destPn, _ = strconv.ParseInt(r.URL.Query().Get("p"), 10, 64)
	} else {
		destPn = 1
	}

	stmtOut, err = db.Prepare(`SELECT a.id, a.username, a.password, a.roleid, a.accessid,  
			b.name, c.pri_rule FROM sp_user a LEFT JOIN sp_role b ON a.roleid = b.id 
			LEFT JOIN sp_access_privilege c  ON a.accessid = c.id ` + con + " GROUP BY a.id ORDER BY a.id DESC LIMIT ?, ?")
	rows, err := stmtOut.Query(cfg.PageSize*(destPn-1), cfg.PageSize)
    defer func() {
        stmtOut.Close()
        rows.Close()
    }()
	if err != nil {
		log.Printf("%s", err)
		http.Redirect(w, r, ERROR_PAGE_NAME, 301)
		return
	}
	for rows.Next() {
		userRelation = &UserRelation{}
		if err := rows.Scan(&userRelation.User.Id, &userRelation.User.UserName, &userRelation.User.Password,
			&userRelation.Role.Id, &userRelation.Access.Id, &userRelation.Role.Name,
			&userRelation.Access.Rule); err != nil {
			log.Printf("%s", err)
			http.Redirect(w, r, ERROR_PAGE_NAME, 301)
			return
		}
		userRelations = append(userRelations, userRelation)
	}
    
	paginator := util.NewPaginator(r, cfg.PageSize, totalN)

	ret := struct {
		Menu      []*SpStatMenu
		Result    []*UserRelation
		Paginator *util.Paginator
		User      *SpStatUser
	}{menu, userRelations, paginator, &user}

	index := strings.LastIndex(path, "/")
	render.HTML(200, path[index+1:], ret)
}



func ViewUserAction(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	redisPool *sexredis.RedisPool, cfg *common.Cfg, session sessions.Session, ms []*SpStatMenu, render render.Render) (int, string) {
	var (
		user SpStatUser
	)
	r.ParseForm()
	id, err := url.QueryUnescape(strings.TrimSpace(r.PostFormValue("id")))
	user.Id, _ = strconv.ParseInt(id, 10, 64)
	stmtOut, err := db.Prepare("SELECT username, password FROM sp_user WHERE id = ?")
	defer stmtOut.Close()
	row := stmtOut.QueryRow(user.Id)
	err = row.Scan(&user.UserName, &user.Password)
	if err != nil {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	if js, err := json.Marshal(user); err != nil {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	} else {
		return http.StatusOK, string(js)
	}
}

func EditUserAction(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	redisPool *sexredis.RedisPool, cfg *common.Cfg, session sessions.Session, ms []*SpStatMenu, render render.Render) (int, string) {
	var (
		user SpStatUser
	)
	r.ParseForm()
	id, err := url.QueryUnescape(strings.TrimSpace(r.PostFormValue("id")))
	user.Id, _ = strconv.ParseInt(id, 10, 64)
	user.UserName, err = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("userName")))
	user.Password, err = url.QueryUnescape(strings.TrimSpace(r.PostFormValue("password")))
    if user.UserName == "" || user.Password == "" {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)        
    }    
	if err != nil {
		log.Printf("post param parse fails %s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	stmtIn, err := db.Prepare("UPDATE sp_user SET username=?, password=? WHERE id = ?")
	defer stmtIn.Close()
	if _, err = stmtIn.Exec(user.UserName, user.Password, user.Id); err != nil {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	js, _ := json.Marshal(Status{"200", "操作成功"})
	return http.StatusOK, string(js)
}
func AddUserAction(r *http.Request, w http.ResponseWriter, db *sql.DB, log *log.Logger,
	redisPool *sexredis.RedisPool, cfg *common.Cfg, session sessions.Session, ms []*SpStatMenu, render render.Render) (int, string) {
	r.ParseForm()
	userName, err := url.QueryUnescape(strings.TrimSpace(r.PostFormValue("userName")))
	password, err := url.QueryUnescape(strings.TrimSpace(r.PostFormValue("password")))
    if userName == "" || password == "" {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)        
    }
	tx, err := db.Begin()
	if err != nil {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	stmtIn, err := tx.Prepare("INSERT INTO sp_user (username, password) VALUES(?, ?)")
	defer stmtIn.Close()

	if err != nil {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}

	result, err := stmtIn.Exec(userName, password)
	if err != nil {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	stmtInAccess, err := tx.Prepare("INSERT INTO sp_access_privilege(pri_group, pri_rule) VALUES(?, ?)")
	defer stmtInAccess.Close()
	if err != nil {
		tx.Rollback()
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	result, err = stmtInAccess.Exec(id, GROUP_PRI_ALLOW)
	if err != nil {
		tx.Rollback()
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	accessId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	stmtInUpdate, err := tx.Prepare("UPDATE sp_user SET accessid = ? WHERE id = ?")
	defer stmtInUpdate.Close()
	if err != nil {
		tx.Rollback()
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	_, err = stmtInUpdate.Exec(accessId, id)

	if err != nil {
		tx.Rollback()
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Printf("%s", err)
		js, _ := json.Marshal(Status{"201", "操作失败"})
		return http.StatusOK, string(js)
	}
	js, _ := json.Marshal(Status{"200", "操作成功"})
	return http.StatusOK, string(js)
}