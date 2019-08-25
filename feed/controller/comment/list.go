package comment

import (
	"encoding/json"
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/service"
	"weibo/lib"
)

type ListReq struct {
	FeedId string `json:"feed_id"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

func (listReq *ListReq) Regular() (ok bool) {
	if listReq == nil {
		return
	}

	if "" == listReq.FeedId {
		return
	}

	ok = true
	return
}

// GetResp 定义输出
type ListResp struct {
	Comment  []*api.Comment `json:"comment"`
	username string         `json:"username"`
}

// @prefilter("Auth")
// @postfilter("Boss")
func (comment *Comment) List(w http.ResponseWriter, r *http.Request) {
	fun := "comment.Comment.List"

	var listReq *ListReq
	if err := json.Unmarshal(comment.ReadBody(r), &listReq); err != nil || !listReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, listReq)
		comment.ReplyFail(w, lib.CodePara)
		return
	}

	commentApi, err := service.NewComment().GetAll(listReq.FeedId)
	if err != nil {
		clog.Error("%s skel.Get err: %v, req: %v", fun, err, nil)
		comment.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &ListResp{
		Comment: commentApi,
	}
	comment.ReplyOk(w, resp)

	return
}
