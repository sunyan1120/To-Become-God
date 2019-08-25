package view

import (
	"encoding/json"
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/service"
	"weibo/lib"
)

type UpdateReq struct {
	Id string `json:"id"`
}

func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil {
		return
	}

	if "" == updateReq.Id {
		return
	}

	ok = true
	return
}

type UpdateResp struct {
}

// @prefilter("Auth")
// @postfilter("Boss")
func (view *View) Add(w http.ResponseWriter, r *http.Request) {
	fun := "view.View.Add"

	var updateReq *UpdateReq
	if err := json.Unmarshal(view.ReadBody(r), &updateReq); err != nil || !updateReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		view.ReplyFail(w, lib.CodePara)
		return
	}

	//看是否存在该feed
	result, err := service.NewFeed().Get(updateReq.Id)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, updateReq)
		view.ReplyFail(w, lib.CodeSrv)
		return
	}
	if result == nil {
		detail := "feed not found"
		view.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	viewApi := api.NewFeed()
	viewApi.ID = updateReq.Id
	viewApi.ViewTime = result.ViewTime
	if err := service.NewView().Add(viewApi); err != nil {
		clog.Error("%s feed.Update err: %v, req: %v", fun, err, updateReq)
		view.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &UpdateResp{}
	view.ReplyOk(w, resp)

	return
}
