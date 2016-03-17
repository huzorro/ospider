package common

import (
	"database/sql"
	"github.com/huzorro/spfactor/sexredis"
	"log"
)

type Ospider struct {
	Log *log.Logger
	P   *sexredis.RedisPool
	Db  *sql.DB
}

type Cfg struct {
	//数据库类型
	Dbtype string `json:"dbtype"`
	//数据库连接uri
	Dburi string `json:"dburi"`
	//页宽
	PageSize int64 `json:"pageSize"`
	//任务队列
	TaskQueueName string `json:"taskQueueName"`
	//结果队列
	ResultQueueName string `json:"resultQueueName"`
	//queue unacked limit
	UnackedLimit int `json:"unackedLimit"`
	//queue consumer num
	NumConsumers int `json:"numConsumers"`
}

type Result struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Processor interface {
	Process(payload string)
}
