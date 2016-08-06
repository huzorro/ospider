package crontab

import (
	"encoding/json"
	"fmt"
	"github.com/adjust/rmq"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler"
	cron "github.com/jakecoffman/cron"
)


type Attack struct {
	common.Ospider
	Cfg
	rmq.Queue
}

func NewAttack() *Attack {
	return &Attack{}
}
func (self *Attack) Handler() {
	c := cron.New()
	subCron := cron.New()
	subCron.Start()
	self.Log.Println(self.FlushCron)
	c.AddFunc(self.FlushCron, func() {
		sqlStr := `select a.id, a.uid, a.url, a.host, a.port, a.method, a.time, 
                   a.powerlevel, a.status, a.cronid, b.name, b.cron 
	               from spider_flood_target a left join spider_crontab b on a.cronid = b.id`

		stmtOut, err := self.Db.Prepare(sqlStr)
		defer stmtOut.Close()

		if err != nil {
			self.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
			return
		}
		var (
            attack handler.FloodTarget
		)
		rows, err := stmtOut.Query()
		defer rows.Close()
		if err != nil {
			self.Log.Printf("stmtOut.Query fails %s", err)
			return
		}
		for rows.Next() {
			if err := rows.Scan(&attack.Id, &attack.Uid, &attack.Url, &attack.Host, 
                                &attack.Port, &attack.Method, &attack.Time,
                                &attack.Powerlevel, &attack.Status, &attack.Crontab.Id, 
                                &attack.Crontab.Name, &attack.Crontab.Cron); err != nil {
				self.Log.Printf("rows.Scan fails %s", err)
				return
			}

			attackJson, _ := json.Marshal(attack)            
            subCron.RemoveJob(fmt.Sprintf("%d", attack.Id))
            if (attack.Status & 1) == 0  {
                continue
            }
			subCron.AddFunc(attack.Crontab.Cron, func() {
				self.Log.Println(attack.Crontab.Cron)
				self.Log.Printf("%s", attackJson)
				//写入redis队列
                if !self.Publish(string(attackJson)) {
                    self.Log.Printf("put in task queue fails")                    
                }
			}, fmt.Sprintf("%d", attack.Id))

		}
	}, "initTask")
	c.Start()
}
