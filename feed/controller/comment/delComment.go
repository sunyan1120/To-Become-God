package comment

import (
	"encoding/json"
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/service"
	"weibo/lib"
)

type DelReq struct {
	FeedId string `json:"feed_id"`
	ID     string `json:"id"`
}

func (delReq *DelReq) Regular() (ok bool) {
	if delReq == nil {
		return
	}

	if "" == delReq.ID {
		return
	}

	ok = true
	return
}

type DelResp struct {
}

// @prefilter("Auth")
// @postfilter("Boss")
func (comment *Comment) Del(w http.ResponseWriter, r *http.Request) {
	fun := "comment.Comment.Del"

	var delReq *DelReq
	if err := json.Unmarshal(comment.ReadBody(r), &delReq); err != nil || !delReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, delReq)
		comment.ReplyFail(w, lib.CodePara)
		return
	}

	if err := service.NewComment().Del(delReq.ID); err != nil {
		clog.Error("%s feed.Del err: %v, req: %v", fun, err, delReq)
		comment.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &DelResp{}
	comment.ReplyOk(w, resp)

	return
}
