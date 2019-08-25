package feed

import (
	"html/template"
	"net/http"
)

// AddReq 定义输入
type TestReq struct {
	Name string `json:"txt"`
}

// Regular 用于参数校验
func (testReq *TestReq) Regular() (ok bool) {
	if testReq == nil {
		return
	}

	if "" == testReq.Name {
		return
	}

	ok = true
	return
}

// @postfilter("Boss")
func (feed *Feed) Test(w http.ResponseWriter, r *http.Request) {
	//fun := "feed.Feed.Add"

	//var testReq *TestReq
	/*if err := json.Unmarshal(feed.ReadBody(r), &testReq); err != nil || !testReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, testReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}*/
	t, _ := template.ParseFiles("././templet/hello.html")
	t.Execute(w, "xxxxxxx")

	/*resp := testReq.Name
	feed.ReplyOk(w, resp)*/

	return
}
