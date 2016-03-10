package processor

import (
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"github.com/huzorro/ospider/crontab"
	"github.com/huzorro/ospider/web/handler"
	// "fmt"
	"bytes"
	"io/ioutil"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

type Cfg struct {
	crontab.Cfg
	CrawlDelay int64  `json:"spiderDelaySecond"`
	MaxVisits  int    `json:"maxVisits"`
	UserAgent  string `json:"userAgent"`
	//并发控制
	ActorNums     int64  `json:"actorNums"`
	LinkQueueName string `json:"linkQueueName"`
}

type ExtLink struct {
	Cfg
	*gocrawl.DefaultExtender
	Attach map[string]interface{}
}

func NewSpiderLink(cfg Cfg, attach map[string]interface{}) (*gocrawl.Crawler, *ExtLink) {
	//attach := make(map[string]interface{})
	ext := &ExtLink{cfg, &gocrawl.DefaultExtender{}, attach}
	// fmt.Println(ext.Attach)
	// Set custom options
	opts := gocrawl.NewOptions(ext)

	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll
	opts.SameHostOnly = true
	opts.MaxVisits = cfg.MaxVisits
	opts.UserAgent = cfg.UserAgent

	return gocrawl.NewCrawlerWithOptions(opts), ext
}

func (e *ExtLink) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	//获取字符集
	var (
		body []byte
		err  error
	)
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		e.Log(gocrawl.LogAll, gocrawl.LogInfo, "response body read fails")
		return nil, true
	}
	encoding, charset, is := charset.DetermineEncoding(body, "utf-8")
	e.Log(gocrawl.LogAll, gocrawl.LogInfo, "response charset is "+charset)
	site := e.Attach["site"].(handler.Site)
	var links []string
	if len(site.Rule.Selector.Href) <= 0 {
		site.Rule.Selector.Href = "a[href]"
	}
	doc.Find(site.Rule.Selector.Section).
		Find(site.Rule.Selector.Href).
		Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		//编码转换
		if !is {
			reader := transform.NewReader(bytes.NewReader([]byte(href)), encoding.NewDecoder())
			if u, err := ioutil.ReadAll(reader); err != nil {
				e.Log(gocrawl.LogAll, gocrawl.LogInfo, "io read all fails"+err.Error())
			} else {
				href = string(u)
			}
		}
		//处理相对路径
		u, _ := url.Parse(href)
		if !u.IsAbs() {
			u.Host = ctx.URL().Host
			u.Scheme = ctx.URL().Scheme
			u.RawQuery = u.Query().Encode()
			href = u.String()
		}
		if len(href) > 0 {
			links = append(links, href)
		}
	})
	e.Attach["links"] = links
	return nil, true
}

func (e *ExtLink) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited
}
