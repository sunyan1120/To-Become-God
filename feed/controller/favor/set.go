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

type SetReq struct {
	Id string `json:"id"`
}

func (setReq *SetReq) Regular() (ok bool) {
	if setReq == nil {
		return
	}

	if "" == setReq.Id {
		return
	}

	ok = true
	return
}

type SetResp struct{}

// @prefilter("Auth")
// @postfilter("Boss")
func (favor *Favor) Set(w http.ResponseWriter, r *http.Request) {
	fun := "favor.Favor.Set"

	var setReq *SetReq
	if err := json.Unmarshal(favor.ReadBody(r), &setReq); err != nil || !setReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, setReq)
		favor.ReplyFail(w, lib.CodePara)
		return
	}
	//取用户id
	uid, err := util.GetUid(favor)
	if err != nil {
		return
	}

	//判断是否存在该feed
	result, err := service.NewFeed().Get(setReq.Id)
	if err != nil {
		clog.Error("%s feed.Get err: %v, req: %v", fun, err, setReq)
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
		clog.Error("%s user.login err: %v, req: %v", fun, err, setReq)
		favor.ReplyFail(w, lib.CodeSrv)
		return
	}
	if user == nil {
		detail := "there is not this user"
		favor.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	favorApi := api.NewFavor()
	favorApi.FeedId = setReq.Id
	favorApi.UserId = uid
	favorApi.UserNick = user.Nick
	if err := service.NewFavor().Set(favorApi); err != nil {
		clog.Error("%s comment.Add err: %v, req: %v", fun, err, setReq)
		favor.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &SetResp{}
	favor.ReplyOk(w, resp)

	return
}
