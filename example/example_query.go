package main

import (
	"github.com/PuerkitoBio/goquery"
    "fmt"
	"strings"
	// "regexp"
	"bytes"
)
func main() {
    var htmls = make([]string, 0)
    doc, _ := goquery.NewDocument("http://www.oschina.net/news/71633/mongodb-3-3-3")
    // doc.
    // Find(`#NewsChannel > div.NewsBody > div > div.NewsEntity > div.Body.NewsContent.TextContent>p,
    // #NewsChannel > div.NewsBody > div > div.NewsEntity > div.Body.NewsContent.TextContent>ul`).
    // Not("p[style]").
    // Each(func(i int, q *goquery.Selection){
    //     fmt.Println(q.Html())
    //     s, _ := q.Html()
    //     htmls = append(htmls, s)        
    // })
    // fmt.Println( strings.Join(htmls, ""))
    // goquery.NewDocumentFromReader(bytes.NewReader
     
    doc.Find("#OSChina_News_71633").Each(func(i int, q *goquery.Selection){
        html, _ := q.Html()
        htmls = append(htmls, html)             
    })
    
    sub, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(strings.Join(htmls, ""))))
    
    sub.Find("img").Each(func(i int, q *goquery.Selection) {
        fmt.Println(q.Attr("src"))
    })
}

