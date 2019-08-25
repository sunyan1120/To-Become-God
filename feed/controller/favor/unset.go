package favor

import (
	"encoding/json"
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/controller/util"
	"weibo/feed/service"
	"weibo/lib"
)

type UnSetReq struct {
	Id string `json:"id"`
}

func (unSetReq *UnSetReq) Regular() (ok bool) {
	if unSetReq == nil {
		return
	}

	if "" == unSetReq.Id {
		return
	}

	ok = true
	return
}

type UnSetResp struct{}

// @prefilter("Auth")
// @postfilter("Boss")
func (favor *Favor) UnSet(w http.ResponseWriter, r *http.Request) {
	fun := "favor.Favor.UnSet"

	var unSetReq *UnSetReq
	if err := json.Unmarshal(favor.ReadBody(r), &unSetReq); err != nil || !unSetReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, unSetReq)
		favor.ReplyFail(w, lib.CodePara)
		return
	}
	//取用户id
	uid, err := util.GetUid(favor)
	if err != nil {
		return
	}

	//判断是否存在该feed
	result, err := service.NewFeed().Get(unSetReq.Id)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, unSetReq)
		favor.ReplyFail(w, lib.CodeSrv)
		return
	}
	if result == nil {
		detail := "there is not this feed"
		favor.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	//根据id去用户
	user, err := service.NewUser().GetById(uid)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, unSetReq)
		favor.ReplyFail(w, lib.CodeSrv)
		return
	}
	if user == nil {
		detail := "there is not this user"
		favor.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	favorApi := api.NewFavor()
	favorApi.FeedId = unSetReq.Id
	favorApi.UserId = uid
	favorApi.UserNick = user.Nick
	if err := service.NewFavor().UnSet(favorApi); err != nil {
		clog.Error("%s comment.Add err: %v, req: %v", fun, err, unSetReq)
		favor.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &UnSetResp{}
	favor.ReplyOk(w, resp)

	return
}
