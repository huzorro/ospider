package processor

import (
	"encoding/json"

	"github.com/adjust/redismq"
	"github.com/gosexy/redis"
	"github.com/huzorro/ospider/web/handler"
)

type TaskContent struct {
	Task
}

func (self TaskContent) Handler() {
	var (
		site     handler.Site
		consumer *redismq.Consumer
		pack     *redismq.Package
		err      error
	)

	linkQueue := redismq.CreateQueue("localhost", "6379", "", 0, self.Cfg.LinkQueueName)
	resultQueue := redismq.CreateQueue("localhost", "6379", "", 0, self.Cfg.ResultQueueName)
	if consumer, err = linkQueue.AddConsumer("consumer"); err != nil {
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
					self.Log.Printf("consumer get Unacked package fails then pack.Fail() %s", err)
					continue
				}
			}

			if err = json.Unmarshal([]byte(pack.Payload), &site); err != nil {
				self.Log.Printf("json Unmarshal fails %s", err)
				continue
			}
			self.Log.Printf("%s", pack.Payload)
			go func() {
				attach := make(map[string]interface{})
				attach["site"] = site
				//查找爬取历史
				var (
					redisClient *redis.Client
					reviewStr   string
					review      handler.Site
				)
				if redisClient, err = self.Ospider.P.Get(); err != nil {
					defer self.Ospider.P.Close(redisClient)
					self.Log.Printf("get redis client fails %s", err)
				}
				if reviewStr, err = redisClient.Get(site.Rule.Url); err != nil {
					self.Log.Printf("not foud history data in review data for key:%s", site.Rule.Url)
					spider, ext := NewSpiderContent(self.Cfg, attach)
					if err := spider.Run(site.Rule.Url); err != nil {
						self.Log.Printf("spider find content fails %s", err)
						site = ext.Attach["site"].(handler.Site)
						if siteStr, err := json.Marshal(site); err != nil {
							self.Log.Printf("site Marshal fails %s", err)
						} else {
							resultQueue.Put(string(siteStr))
							//写入redis－历史记录
							if redisClient, err := self.Ospider.P.Get(); err != nil {
								defer self.Ospider.P.Close(redisClient)
								self.Log.Printf("get redis client fails %s", err)
							} else {
								redisClient.Set(site.Rule.Url, string(siteStr))
								redisClient.Expire(site.Rule.Url, 60*60*48)
							}
						}
					}
				} else {
					if err = json.Unmarshal([]byte(reviewStr), &review); err != nil {
						self.Log.Printf("json Unmarshal fails %s", err)
					} else {
						if review.Id != site.Id {
							site.Rule.Selector.Title = review.Rule.Selector.Title
							site.Rule.Selector.Content = review.Rule.Selector.Content
							if siteStr, err := json.Marshal(site); err != nil {
								self.Log.Printf("site Marshal fails %s", err)
							} else {
								resultQueue.Put(string(siteStr))
							}
						}
					}
				}
				pack.Ack()
				<-sem
			}()

		}
	}()
}
