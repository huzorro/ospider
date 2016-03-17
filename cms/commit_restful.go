package cms

import (
	"encoding/json"
	"net/http"
	"strings"
    "time"
	"github.com/huzorro/ospider/common"
	"github.com/huzorro/ospider/web/handler"
)

type CommitRestful struct {
	common.Ospider
}

type Document struct {
    Uid int64 `json:"uid"`
    Title string `json:"title"`
    Category int64 `json:"category_id"`
    Group int64 `json:"group_id"`
    Model int64 `json:"model_id"`
    Position int64 `json:"position"`
    Display int64 `json:"display"`
    Status int64 `json:"status"`
    Create int64 `json:"create_time"`
    Update int64 `json:"update_time"`
}

type Article struct {
    Id int64 `json:"id"`
    Content string `json:"content"`
}

type DocumentArticle struct {
    Document    `json:"document"`
    Article     `json:"article"`
}
func NewCommitRestful(co common.Ospider) *CommitRestful {
	return &CommitRestful{co}
}

func (self *CommitRestful) Process(payload string) {
	var (
		site handler.Site
        documentArticle DocumentArticle
        resultJson []byte
        err error
	)
	if err = json.Unmarshal([]byte(payload), &site); err != nil {
		self.Log.Printf("json Unmarshal fails %s", err)
		return
	}
    //
    documentArticle.Article.Content = site.Rule.Selector.Content
    documentArticle.Document.Category = site.DocumentSet.Category
    documentArticle.Document.Display = site.DocumentSet.Display
    documentArticle.Document.Group = site.DocumentSet.GroupId
    documentArticle.Document.Model = site.DocumentSet.ModelId
    documentArticle.Document.Position = site.DocumentSet.Position
    documentArticle.Document.Status = site.DocumentSet.Check
    documentArticle.Document.Title = site.Rule.Selector.Title
    documentArticle.Document.Uid = site.Uid
    documentArticle.Document.Create = time.Now().Unix()
    documentArticle.Document.Update = time.Now().Unix()
    
    if resultJson, err = json.Marshal(documentArticle); err != nil {
        self.Log.Printf("json Marshal fails %s", err)
        return
    }             
	client := &http.Client{}
	req, _ := http.NewRequest("POST", site.Url + "/api/document/article", strings.NewReader(string(resultJson)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
    defer resp.Body.Close()
    self.Log.Printf("commit resuful:[%d]:%s", resp.StatusCode, site.Url + "/document/article") 
	// {"document":{"uid:":1,"title":"test112","category_id":42,"group_id":0,"model_id":2,"position":4,"display":1,"status":1},
	// "article":{"id":,"content":"adfdsafsafsd"}}
}
