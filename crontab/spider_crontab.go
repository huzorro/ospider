package crontab

import (
    
    "github.com/huzorro/ospider/common"
	"github.com/robfig/cron"
	"github.com/huzorro/ospider/web/handler"
	"encoding/json"
	"github.com/adjust/redismq"
)

type Cfg struct {
    common.Cfg
    FlushCron string `json:"flushCron"`            
}
 
type Task struct {
    common.Ospider
    Cfg
    *redismq.Queue     
}
func New() *Task {
    return &Task{}
}
func (self *Task) Handler() {  
	c := cron.New()
    subCron := cron.New()
    self.Log.Println(self.FlushCron)
	c.AddFunc(self.FlushCron, func() {
        sqlStr := `select a.id, a.uid, a.name, a.url, a.document_set, 
                    a.ruleid, b.name, b.spiderid,b.url, b.rule_json,  
                    a.cronid, c.name, c.cron, d.id, d.name, d.queue_name_json 
                    from spider_site a left join spider_rule b on a.ruleid = b.id 
                    left join spider_crontab c on a.cronid = c.id 
                    left join spider_manager d on b.spiderid = d.id where a.status = 1`
                    
        stmtOut, err := self.Db.Prepare(sqlStr)
        defer stmtOut.Close()
        
        if err != nil {
            self.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
            return
        }
        var (
            site handler.Site 
        )
        rows, err := stmtOut.Query()
        defer rows.Close()
        if err != nil {
            self.Log.Printf("stmtOut.Query fails %s", err)
            return            
        }
        for rows.Next() {
            if err := rows.Scan(&site.Id, &site.Uid, &site.Name, &site.Url, &site.DocumentSetStr, &site.Rule.Id,
                                &site.Rule.Name, &site.Rule.SpiderId, &site.Rule.Url, &site.Rule.SelectorStr,
                                &site.Crontab.Id, &site.Crontab.Name, &site.Crontab.Cron, &site.Rule.Manager.Id,
                                &site.Rule.Manager.Name, &site.Rule.Manager.QueueNameStr); err != nil {
                self.Log.Printf("rows.Scan fails %s", err)
                return                                    
            }
            if err := json.Unmarshal([]byte(site.DocumentSetStr), &site.DocumentSet); err != nil {
                self.Log.Printf("Unmarshal DocumentSetStr to DocumentSet fails %s", err)
                return
            }
            if err := json.Unmarshal([]byte(site.Rule.SelectorStr), &site.Rule.Selector); err != nil {
                self.Log.Printf("Unmarshal SelectorStr to Selector fails %s", err)
                return
            }
            if err := json.Unmarshal([]byte(site.Rule.Manager.QueueNameStr), &site.Rule.Manager.QueueName); err != nil {
                self.Log.Printf("Unmarshal SelectorStr to Selector fails %s", err)
                return                
            }
            
            siteJson, _ := json.Marshal(site)
            self.Log.Printf("%s", siteJson)            
            subCron.Stop()
            subCron.AddFunc(site.Crontab.Cron, func() {
                self.Log.Println(site.Crontab.Cron)
                self.Log.Printf("%s", siteJson)
                //写入redis队列
                if err := self.Put(string(siteJson)); err != nil {
                    self.Log.Printf("put in task queue fails %s", err)
                }                
            })
            subCron.Start()
        }               
	}) 
	c.Start()
}