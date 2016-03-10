package processor



import (
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler"
	"github.com/adjust/redismq"
	"encoding/json"
)

type Task struct {
    Cfg
    common.Ospider
    *redismq.Queue          
}
func (self Task) Handler() {
    var(
        site handler.Site
        consumer *redismq.Consumer
        pack *redismq.Package
        err error
    )
    
    if consumer, err = self.Queue.AddConsumer("consumer"); err != nil {
        self.Log.Printf("Queue add consumer fails %s", err)
        return
    }
    linkQueue := redismq.CreateQueue("localhost", "6379", "", 0, self.Cfg.LinkQueueName)
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
                attach := make(map[string]interface{})
                attach["site"] = site
                spider, ext := NewSpiderLink(self.Cfg, attach)                
                if err := spider.Run(site.Rule.Url); err != nil { 
                    self.Log.Printf("spider find link fails %s", err) 
                    self.Log.Println(ext.Attach)  
                    links := ext.Attach["links"].([]string)
                    for _, v := range links {
                        site.Rule.Url = v
                        if siteStr, err := json.Marshal(site); err != nil {
                            self.Log.Printf("json Marshal fails %s", err)
                        } else {                            
                            linkQueue.Put(string(siteStr))
                        }
                        
                    }                                      
                } 
                pack.Ack()
                <- sem   
            }()
            
        }
    }()
}