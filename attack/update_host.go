package attack

import (
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler" 
	"encoding/json"
    "github.com/huzorro/ospider/util"
    "net"
)

type UpdateHost struct {
    cfg common.Ospider
}
func NewUpdateHost(co common.Ospider) *UpdateHost {
    return &UpdateHost{cfg:co}
}

func (self *UpdateHost) Process(payload string) {
    var (
        attack handler.FloodTarget
    )
    self.cfg.Log.Println("update host...")
    if err := json.Unmarshal([]byte(payload), &attack); err != nil {
        self.cfg.Log.Printf("json Unmarshal fails %s", err)
        return
    }
    
    sqlStr := `update spider_flood_target set host = ? where url = ? and status = 1`
    
    stmtIn, err := self.cfg.Db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        self.cfg.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return
    } 

    ip, err := util.LookupHost(attack.Url)
    if err != nil {
        self.cfg.Log.Println("lookup host fails for api") 
        ips, err := net.LookupIP(attack.Url)
        if  err == nil {
            if ips[0].String() != "127.0.0.1" {
                attack.Host = ips[0].String()
            }
        }                
    } else {                
        attack.Host = ip
    }    
    self.cfg.Log.Printf("lookup host %s", attack.Host)    
    _, err = stmtIn.Exec(attack.Host, attack.Url)
    
    if err != nil {
        self.cfg.Log.Printf("update flood target host fails %s", err)
        return
    } 
                     
}
