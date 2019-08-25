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

type SubUserReq struct {
	Need_sub_uid string `json:"need_sub_uid"`
}

func (subUserReq *SubUserReq) Regular() (ok bool) {
	if subUserReq == nil {
		return
	}

	if "" == subUserReq.Need_sub_uid {
		return
	}

	ok = true
	return
}

type SubUserResp struct{}

// @prefilter("Auth")
// @postfilter("Boss")
func (user *User) SubUser(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.sub_user"

	var subUserReq *SubUserReq
	if err := json.Unmarshal(user.ReadBody(r), &subUserReq); err != nil || !subUserReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, subUserReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}
	//取用户id
	uid, err := util.GetUid(user)
	if err != nil {
		return
	}

	//判断是否关注了该用户
	result, _ := service.NewAttention().IsExistSub(uid, subUserReq.Need_sub_uid)
	if result != 0 {
		detail := "已关注过该用户"
		clog.Error(fun, detail)
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	//判断是否存在该用户
	him, err := service.NewUser().GetById(subUserReq.Need_sub_uid)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, subUserReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	if him == nil {
		detail := "there is not this user"
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}
	//关注
	attentionApi := api.NewAttention()
	attentionApi.AttId = subUserReq.Need_sub_uid
	attentionApi.UserId = uid

	if err := service.NewAttention().AddAttention(attentionApi); err != nil {
		clog.Error("%s user.addAtt err: %v, req: %v", fun, err, subUserReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &SubUserResp{}
	user.ReplyOk(w, resp)

	return

}
