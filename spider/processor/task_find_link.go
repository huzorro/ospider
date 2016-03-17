package processor

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/adjust/rmq"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler"
)

type Task struct {
	Cfg
	common.Ospider
	rmq.Connection
}
type Consumer struct {
	name   string
	count  int
	before time.Time
	task   Task
}

func NewConsumer(tag int, task Task) *Consumer {
	return &Consumer{
		name:   fmt.Sprintf("consumer%d", tag),
		count:  0,
		before: time.Now(),
		task:   task,
	}
}
func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	var (
		site      handler.Site
		linkQueue rmq.Queue
		err       error
	)
	consumer.count++
    linkQueue = consumer.task.Connection.OpenQueue(consumer.task.Cfg.LinkQueueName)
	consumer.task.Log.Printf("%s consumed %d %s", consumer.name, consumer.count, delivery.Payload())
	if err = json.Unmarshal([]byte(delivery.Payload()), &site); err != nil {
		consumer.task.Log.Printf("json Unmarshal fails %s", err)
		delivery.Ack()
		return
	}
	attach := make(map[string]interface{})
	attach["site"] = site
	spider, ext := NewSpiderLink(consumer.task.Cfg, attach)
	if err := spider.Run(site.Rule.Url); err != nil {
		links := ext.Attach["links"].([]string)
		for _, v := range links {
			site.Rule.Url = v
			if siteStr, err := json.Marshal(site); err != nil {
				consumer.task.Log.Printf("json Marshal fails %s", err)
			} else {
				linkQueue.Publish(string(siteStr))
			}

		}
	}
	delivery.Ack()
}

func (self Task) Handler() {
	var (
		taskQueue rmq.Queue
	)
	taskQueue = self.Connection.OpenQueue(self.Cfg.TaskQueueName)
	taskQueue.StartConsuming(self.UnackedLimit, 500*time.Millisecond)

	for i := 0; i < self.NumConsumers; i++ {
		name := fmt.Sprintf("consumer %d", i)
		taskQueue.AddConsumer(name, NewConsumer(i, self))
	}
}
