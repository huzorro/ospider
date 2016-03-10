package main

import (
	"github.com/PuerkitoBio/goquery"
    "fmt"
	"strings"
)
func main() {
    var htmls = make([]string, 0)
    doc, _ := goquery.NewDocument("http://www.oschina.net/news/71395/sonarqube-5-4")
    doc.
    Find(`#NewsChannel > div.NewsBody > div > div.NewsEntity > div.Body.NewsContent.TextContent>p,
    #NewsChannel > div.NewsBody > div > div.NewsEntity > div.Body.NewsContent.TextContent>ul`).
    Not("p[style]").
    Each(func(i int, q *goquery.Selection){
        fmt.Println(q.Html())
        s, _ := q.Html()
        htmls = append(htmls, s)        
    })
    fmt.Println( strings.Join(htmls, ""))
}

