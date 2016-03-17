package processor

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/adjust/rmq"
	"github.com/gosexy/redis"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler"
)

type TaskContent struct {
	Cfg
	common.Ospider
	rmq.Connection
}
type ConsumerContent struct {
	name    string
	count   int
	before  time.Time
	content TaskContent
}

func NewConsumerLink(tag int, task TaskContent) *ConsumerContent {
	return &ConsumerContent{
		name:    fmt.Sprintf("consumer%d", tag),
		count:   0,
		before:  time.Now(),
		content: task,
	}
}
func (consumer *ConsumerContent) Consume(delivery rmq.Delivery) {
	var (
		site        handler.Site
		resultQueue rmq.Queue
		err         error
	)
	consumer.count++
	resultQueue = consumer.content.Connection.OpenQueue(consumer.content.Cfg.ResultQueueName)
	consumer.content.Log.Printf("%s consumed %d %s", consumer.name, consumer.count, delivery.Payload())
	if err = json.Unmarshal([]byte(delivery.Payload()), &site); err != nil {
		consumer.content.Log.Printf("json Unmarshal fails %s", err)
		delivery.Ack()
		return
	}
	attach := make(map[string]interface{})
	attach["site"] = site
	//查找爬取历史
	var (
		redisClient *redis.Client
		reviewStr   string
		review      handler.Site
	)
	if redisClient, err = consumer.content.Ospider.P.Get(); err != nil {
		defer consumer.content.Ospider.P.Close(redisClient)
		consumer.content.Log.Printf("get redis client fails %s", err)
	}
	if reviewStr, err = redisClient.Get(site.Rule.Url); err != nil {
		consumer.content.Log.Printf("not foud history data in review data for key:%s", site.Rule.Url)
		spider, ext := NewSpiderContent(consumer.content.Cfg, attach)
		if err := spider.Run(site.Rule.Url); err != nil {
			consumer.content.Log.Printf("spider find content fails %s", err)
			site = ext.Attach["site"].(handler.Site)
			if siteStr, err := json.Marshal(site); err != nil {
				consumer.content.Log.Printf("site Marshal fails %s", err)
			} else {
				resultQueue.Publish(string(siteStr))
				//写入redis－历史记录
				if redisClient, err := consumer.content.Ospider.P.Get(); err != nil {
					defer consumer.content.Ospider.P.Close(redisClient)
					consumer.content.Log.Printf("get redis client fails %s", err)
				} else {
					redisClient.Set(site.Rule.Url, string(siteStr))
					redisClient.Expire(site.Rule.Url, 60*60*48)
				}
			}
		}
	} else {
		if err = json.Unmarshal([]byte(reviewStr), &review); err != nil {
			consumer.content.Log.Printf("json Unmarshal fails %s", err)
		} else {
			if review.Id != site.Id || review.Url != site.Url {
				site.Rule.Selector.Title = review.Rule.Selector.Title
				site.Rule.Selector.Content = review.Rule.Selector.Content
				if siteStr, err := json.Marshal(site); err != nil {
					consumer.content.Log.Printf("site Marshal fails %s", err)
				} else {
					resultQueue.Publish(string(siteStr))
				}
			}
		}
	}
    delivery.Ack()
}
func (self TaskContent) Handler() {
	var (
		linkQueue rmq.Queue
	)
	linkQueue = self.Connection.OpenQueue(self.Cfg.LinkQueueName)
	linkQueue.StartConsuming(self.UnackedLimit, 500*time.Millisecond)

	for i := 0; i < self.NumConsumers; i++ {
		name := fmt.Sprintf("consumer %d", i)
		linkQueue.AddConsumer(name, NewConsumerLink(i, self))
	}
}

