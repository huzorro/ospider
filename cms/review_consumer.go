package cms

import (    
	"fmt"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/spider/processor"    
	"github.com/adjust/redismq"
)
type Consumer struct {
	name   string
    cfg processor.Cfg
    co common.Ospider
    processors []common.Processor
    
}
func NewConsumer(tag int, cfg processor.Cfg, co common.Ospider) *Consumer {
	c := &Consumer{
		name:   fmt.Sprintf("consumer%d", tag),
        cfg:cfg,
        co:co,
	}
    c.processors = make([]common.Processor, 0)
    return c
}

func (consumer *Consumer) AddProcessor(p common.Processor) *Consumer {
    consumer.processors = append(consumer.processors, p) 
    return consumer 
}
func (consumer *Consumer) Consume(pack *redismq.Package) {    
    for _, p := range consumer.processors {
        p.Process(pack.Payload)
    }
    // pack.Ack()
}