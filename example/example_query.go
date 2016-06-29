package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	// "regexp"
	"bytes"
	"encoding/binary"
	"net"
    "github.com/huzorro/ospider/util"
)

func netName() {
	addr, _ := net.InterfaceAddrs()
	fmt.Println(addr) //[127.0.0.1/8 10.236.15.24/24 ::1/128 fe80::3617:ebff:febe:f123/64],本地地址,ipv4和ipv6地址,这些信息可通过ifconfig命令看到
	interfaces, _ := net.Interfaces()
	fmt.Println(interfaces) //[{1 65536 lo  up|loopback} {2 1500 eth0 34:17:eb:be:f1:23 up|broadcast|multicast}] 类型:MTU(最大传输单元),网络接口名,支持状态
	hp := net.JoinHostPort("127.0.0.1", "8080")
	fmt.Println(hp) //127.0.0.1:8080,根据ip和端口组成一个addr字符串表示
	lt, _ := net.LookupAddr("127.0.0.1")
	fmt.Println(lt) //[localhost],根据地址查找到改地址的一个映射列表
	cname, _ := net.LookupCNAME("www.wxbsj.com")
	fmt.Println(cname) //www.a.shifen.com,查找规范的dns主机名字
	host, _ := net.LookupHost("www.wxbsj.com")
	fmt.Println(host) //[111.13.100.92 111.13.100.91],查找给定域名的host名称
	ip, _ := net.LookupIP("www.hndydb.com")
	fmt.Println(ip) //[111.13.100.92 111.13.100.91],查找给定域名的ip地址,可通过nslookup www.baidu.com进行查找操作.
	fmt.Println(ip[0].String())
    ipp, err := util.LookupHost("www.juanl.net")
    
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("ip:%s ``", ipp)
}

func pack2unpack() {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint32(buf[:4], uint32(9000))
	binary.BigEndian.PutUint32(buf[4:], uint32(9000))


	fmt.Println(binary.BigEndian.Uint32(buf[:4]))
	fmt.Println(binary.BigEndian.Uint32(buf[4:]))
	fmt.Println(binary.BigEndian.Uint64(buf))
    
    binary.BigEndian.PutUint32(buf[:4], uint32(2500))
	binary.BigEndian.PutUint32(buf[4:], uint32(1))
    
    fmt.Println(binary.BigEndian.Uint32(buf[:4]))
	fmt.Println(binary.BigEndian.Uint32(buf[4:]))
	fmt.Println(binary.BigEndian.Uint64(buf))
    //19327352841000
    
    var unbuf = make([]byte, 8)
    binary.BigEndian.PutUint64(unbuf, uint64(38654705670000))
    fmt.Println(binary.BigEndian.Uint32(unbuf[:4]))
	fmt.Println(binary.BigEndian.Uint32(unbuf[4:]))
    
}
func main() {
	var htmls = make([]string, 0)
	doc, _ := goquery.NewDocument("http://www.oschina.net/news/71633/mongodb-3-3-3")
	// doc.
	// Find(`#NewsChannel > div.NewsBody > div > div.NewsEntity > div.Body.NewsContent.TextContent>p,
	// #NewsChannel > div.NewsBody > div > div.NewsEntity > div.Body.New
	// Each(func(i int, q *goquery.Selection){
	//     fmt.Println(q.Html())
	//     s, _ := q.Html()
	//     htmls = append(htmls, s)
	// })
	// fmt.Println( strings.Join(htmls, ""))
	// goquery.NewDocumentFromReader(bytes.NewReader

	doc.Find("#OSChina_News_71633").Each(func(i int, q *goquery.Selection) {
		html, _ := q.Html()
		htmls = append(htmls, html)
	})

	sub, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(strings.Join(htmls, ""))))

	sub.Find("img").Each(func(i int, q *goquery.Selection) {
		fmt.Println(q.Attr("src"))
	})

	netName()

	pack2unpack()
}
