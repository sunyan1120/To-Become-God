package user

import (
	"encoding/json"
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/controller/util"
	"weibo/feed/service"
	"weibo/lib"
)

type UnSubReq struct {
	Need_unsub_uid string `json:"need_unsub_uid"`
}

func (unSubReq *UnSubReq) Regular() (ok bool) {
	if unSubReq == nil {
		return
	}

	if "" == unSubReq.Need_unsub_uid {
		return
	}

	ok = true
	return
}

type UnSubResp struct{}

// @prefilter("Auth")
// @postfilter("Boss")
func (user *User) UnsubUser(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.unsub_user"

	var unSubReq *UnSubReq
	if err := json.Unmarshal(user.ReadBody(r), &unSubReq); err != nil || !unSubReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, unSubReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}
	//取用户id
	uid, err := util.GetUid(user)
	if err != nil {
		return
	}
	//判断是否关注了该用户
	result, _ := service.NewAttention().IsExistSub(uid, unSubReq.Need_unsub_uid)
	if result == 0 {
		detail := "未关注过该用户"
		clog.Error(fun, detail)
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	//取消关注
	attentionApi := api.NewAttention()
	attentionApi.AttId = unSubReq.Need_unsub_uid
	attentionApi.UserId = uid

	if err := service.NewAttention().DelAttention(attentionApi); err != nil {
		clog.Error("%s user.addAtt err: %v, req: %v", fun, err, unSubReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &UnSubResp{}
	user.ReplyOk(w, resp)

	return
}
