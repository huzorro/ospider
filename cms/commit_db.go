package cms

import (
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler"    
	"encoding/json"
)

type CommitMysql struct {
    common.Ospider                
}
func NewCommitMysql(co common.Ospider) *CommitMysql {
    return &CommitMysql{co}
}

func (self *CommitMysql) Process(payload string) {
    var (
        site handler.Site
    )
    if err := json.Unmarshal([]byte(payload), &site); err != nil {
        self.Log.Printf("json Unmarshal fails %s", err)
        return
    }
    sqlStr := `insert into spider_history (uid, siteid, url, result_json) 
                values(?, ?, ?, ?)`
    stmtIn, err := self.Db.Prepare(sqlStr)
    defer stmtIn.Close()
    if err != nil {
        self.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return
    }
    if _, err := stmtIn.Exec(site.Uid, site.Id, site.Rule.Url, payload); err != nil {
        self.Log.Printf("stmtIn.Exec() fails %s", err)
        return
    }
}
