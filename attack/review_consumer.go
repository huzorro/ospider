package attack

import (
	"fmt"
	"time"

	"github.com/adjust/rmq"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/spider/processor"
)

type Consumer struct {
	name   string
	count  int
	before time.Time
	task   *Task
}

type Task struct {
	Cfg processor.Cfg
	common.Ospider
	rmq.Connection
	Processors []common.Processor
}

func NewConsumer(tag int, task *Task) *Consumer {
	c := &Consumer{
		name:   fmt.Sprintf("consumer%d", tag),
		count:  0,
		before: time.Now(),
		task:   task,
	}
	return c
}

func (self *Task) AddProcessor(p common.Processor) *Task {
	self.Processors = append(self.Processors, p)
	return self
}
func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	consumer.count++
    elem := delivery.Payload() 
	consumer.task.Log.Printf("%s consumed %d %s", consumer.name, consumer.count, elem)
	for _, p := range consumer.task.Processors {
		p.Process(elem)
	}
	delivery.Ack()
}

func (self *Task) Handler() {
	var (
		resultQueue rmq.Queue
	)
	resultQueue = self.Connection.OpenQueue(self.Cfg.AttackQueueName)
	resultQueue.StartConsuming(self.Cfg.UnackedLimit, 500*time.Millisecond)

	for i := 0; i < self.Cfg.NumConsumers; i++ {
		name := fmt.Sprintf("consumer %d", i)
		resultQueue.AddConsumer(name, NewConsumer(i, self))
	}
}
