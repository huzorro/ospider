package cms

import (
	"encoding/json"

	"github.com/adjust/redismq"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/spider/processor"
	"github.com/huzorro/ospider/web/handler"
)

type ResultService struct {
	Cfg processor.Cfg
	common.Ospider
	*redismq.Queue
    Consumer *Consumer
}

func NewResultService(cfg processor.Cfg, co common.Ospider, q *redismq.Queue, c *Consumer) *ResultService {
    r := &ResultService{cfg, co, q, c}
    return r 
}
func (self ResultService) Handler() {
	var (
		site     handler.Site
		consumer *redismq.Consumer
		pack     *redismq.Package
		err      error
	)

	if consumer, err = self.Queue.AddConsumer("consumer_cms_1"); err != nil {
		self.Log.Printf("Queue add consumer fails %s", err)
		return
	}
	go func() {
		sem := make(chan interface{}, self.Cfg.ActorNums)
		defer close(sem)
		for {
			sem <- 0
			if pack, err = consumer.Get(); err != nil {
				self.Log.Printf("consumer get  package fails then pack.Fail() %s", err)
				if pack, err = consumer.GetUnacked(); err != nil {
					self.Log.Printf("consumer get GetUnacked package fails then pack.Fail() %s", err)
					continue
				}
			}

			if err = json.Unmarshal([]byte(pack.Payload), &site); err != nil {
				self.Log.Printf("json Unmarshal fails %s", err)
				continue
			}
			self.Log.Printf("%s", pack.Payload)
            go func() {
                self.Consumer.Consume(pack)
                <-sem
            }()
			// pack.Ack()
		}
	}()
}
