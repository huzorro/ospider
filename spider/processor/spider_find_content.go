package processor

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"github.com/huzorro/ospider/web/handler"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"strings"
)

type ExtContent struct {
	Cfg
	*gocrawl.DefaultExtender
	Attach map[string]interface{}
}

func NewSpiderContent(cfg Cfg, attach map[string]interface{}) (*gocrawl.Crawler, *ExtContent) {
	//attach := make(map[string]interface{})
	ext := &ExtContent{cfg, &gocrawl.DefaultExtender{}, attach}
	// Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogAll
	opts.SameHostOnly = true
	opts.MaxVisits = cfg.MaxVisits
	opts.UserAgent = cfg.UserAgent

	return gocrawl.NewCrawlerWithOptions(opts), ext
}

func (e *ExtContent) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	//获取字符集
	var (
		body []byte
		err  error
        htmls = make([]string, 0)
	)
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		e.Log(gocrawl.LogAll, gocrawl.LogInfo, "response body read fails")
		return nil, true
	}
	encoding, charset, is := charset.DetermineEncoding(body, "utf-8")
    e.Log(gocrawl.LogAll, gocrawl.LogInfo, "response charset is " + charset)    
	site := e.Attach["site"].(handler.Site)
	site.Rule.Selector.Title = doc.Find(site.Rule.Selector.Title).Text()
	//编码转换
	if !is {
		reader := transform.NewReader(bytes.NewReader([]byte(site.Rule.Selector.Title)), encoding.NewDecoder())
		if u, err := ioutil.ReadAll(reader); err != nil {
			e.Log(gocrawl.LogAll, gocrawl.LogInfo, "io read all fails"+err.Error())
		} else {
			site.Rule.Selector.Title = string(u)
		}
	}
     
    if len(site.Rule.Selector.Filter) > 0 {
       doc.Find(site.Rule.Selector.Content).
       Not(site.Rule.Selector.Filter).
       Each(func(i int, q *goquery.Selection){
            shtml, _ := q.Html()
            htmls = append(htmls, shtml)                                   
       })       
    } else {
        doc.Find(site.Rule.Selector.Content).
        Each(func(i int, q *goquery.Selection){
            shtml, _ := q.Html()
            htmls = append(htmls, shtml)                                   
       })         
    }
    site.Rule.Selector.Content = strings.Join(htmls, "<br/>")
    //编码转换
	if !is {
		reader := transform.NewReader(bytes.NewReader([]byte(site.Rule.Selector.Content)), encoding.NewDecoder())
		if u, err := ioutil.ReadAll(reader); err != nil {
			e.Log(gocrawl.LogAll, gocrawl.LogInfo, "io read all fails"+err.Error())
		} else {
			site.Rule.Selector.Content = string(u)
		}
	}
	e.Attach["site"] = site
	return nil, true
}

func (e *ExtContent) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	return !isVisited
}
