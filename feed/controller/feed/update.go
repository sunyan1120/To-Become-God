package feed

import (
	"encoding/json"
	"net/http"

	"weibo/feed/api"
	"weibo/feed/service"
	"weibo/lib"

	"github.com/simplejia/clog"
)

type UpdateReq struct {
	Id  string `json:"id" bson:"_id"`
	Txt string `json:"txt" bson:"_txt"`
}

func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil {
		return
	}

	if "" == updateReq.Txt || "" == updateReq.Id {
		return
	}

	ok = true
	return
}

type UpdateResp struct {
}

// @prefilter("Auth")
// @postfilter("Boss")
func (feed *Feed) Update(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Update"

	var updateReq *UpdateReq
	if err := json.Unmarshal(feed.ReadBody(r), &updateReq); err != nil || !updateReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}
	result, err := service.NewFeed().Get(updateReq.Id)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}
	if result == nil {
		detail := "feed not found"
		feed.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}
	feedApi := api.NewFeed()
	feedApi.ID = updateReq.Id
	feedApi.Txt = updateReq.Txt
	if err := service.NewFeed().Update(feedApi); err != nil {
		clog.Error("%s feed.Update err: %v, req: %v", fun, err, updateReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &UpdateResp{}
	feed.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.UPDATE, nil)

	return
}
