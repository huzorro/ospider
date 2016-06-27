package crontab

import (
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler"
	cron "github.com/jakecoffman/cron"
    "encoding/binary"
)


type Rest struct {
	common.Ospider
	Cfg
}

func NewRest() *Rest {
	return &Rest{}
}
func (self *Rest) Handler() {
	c := cron.New()
	subCron := cron.New()
	subCron.Start()
	self.Log.Println(self.RestCron)
	c.AddFunc(self.RestCron, func() {
    var (
        api handler.FloodApi       
        powerlevel uint32
        time uint32
    )
    self.Log.Println("rest powerlevel and time...")
    sqlStr := `select id, name, api, powerlevel, time 
                from spider_flood_api where status = 1`            
    stmtOut, err := self.Db.Prepare(sqlStr)
    defer stmtOut.Close()
    if err != nil {
        self.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return
    }  
    
     
    rows, err := stmtOut.Query()
    if err != nil {
        self.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
        return        
    }
    for rows.Next() {
        err = rows.Scan(&api.Id, &api.Name, &api.Api, &api.Powerlevel, &api.Time)
        if err != nil {
            self.Log.Printf("rows.Scan (%s) fails %s", sqlStr, err)
            return             
        }
        var powerlevelBuf = make([]byte, 8)
        binary.BigEndian.PutUint64(powerlevelBuf, uint64(api.Powerlevel))
        powerlevel = binary.BigEndian.Uint32(powerlevelBuf[:4])
        binary.BigEndian.Uint32(powerlevelBuf[4:])
        
        var timeBuf = make([]byte, 8)
        binary.BigEndian.PutUint64(timeBuf, uint64(api.Time))
        time = binary.BigEndian.Uint32(timeBuf[:4])
        binary.BigEndian.Uint32(timeBuf[4:])
        
        sqlStr = `update spider_flood_api set time = ?, powerlevel = ? where id = ?`
        
        stmtIn, err := self.Db.Prepare(sqlStr)
        defer stmtIn.Close()
        if err != nil {
            self.Log.Printf("db.Prepare(%s) fails %s", sqlStr, err)
            return
        } 

        binary.BigEndian.PutUint32(powerlevelBuf[:4], powerlevel)
        binary.BigEndian.PutUint32(powerlevelBuf[4:], powerlevel) 
        
        

        binary.BigEndian.PutUint32(timeBuf[:4], time)
        binary.BigEndian.PutUint32(timeBuf[4:], time)
        
        newPowerlevel := binary.BigEndian.Uint64(powerlevelBuf)
        newTime := binary.BigEndian.Uint64(timeBuf)
        
        self.Log.Printf("powerlevel:%d-%d time:%d-%d", powerlevel, newPowerlevel, time, newTime )
        _, err = stmtIn.Exec(newTime, newPowerlevel, api.Id)
        
        if err != nil {
            self.Log.Printf("update flood api fails %s", err)
            return
        }
    }}, "initTask")
	c.Start()
}
