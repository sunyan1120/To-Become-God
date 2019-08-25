package comment

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/controller/util"
	"weibo/feed/service"
	"weibo/lib"
)

type AddReq struct {
	FeedId     string `json:"feed_id"`
	CommentTxt string `json:"comment_txt"`
}

func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	if "" == addReq.FeedId || "" == addReq.CommentTxt {
		return
	}

	ok = true
	return
}

type AddResp struct{}

// @prefilter("Auth")
// @postfilter("Boss")
func (comment *Comment) Add(w http.ResponseWriter, r *http.Request) {
	fun := "comment.Comment.Add"

	var addReq *AddReq
	if err := json.Unmarshal(comment.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		comment.ReplyFail(w, lib.CodePara)
		return
	}

	uid, err := util.GetUid(comment)
	if err != nil {
		return
	}

	//判断是否存在该feed
	result, err := service.NewFeed().Get(addReq.FeedId)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, addReq)
		comment.ReplyFail(w, lib.CodeSrv)
		return
	}
	if result == nil {
		detail := "there is not this feed"
		comment.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}
	//根据id去用户
	user, err := service.NewUser().GetById(uid)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, addReq)
		comment.ReplyFail(w, lib.CodeSrv)
		return
	}
	if user == nil {
		detail := "there is not this user"
		comment.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}
	//添加
	commentApi := api.NewComment()
	commentApi.FeedId = addReq.FeedId
	commentApi.ID = uuid.Must(uuid.NewV4()).String()
	commentApi.CommentTxt = addReq.CommentTxt
	commentApi.PostUserId = uid
	commentApi.UserNick = user.Nick
	if err := service.NewComment().AddComment(commentApi); err != nil {
		clog.Error("%s comment.Add err: %v, req: %v", fun, err, addReq)
		comment.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddResp{}
	comment.ReplyOk(w, resp)

	return
}
