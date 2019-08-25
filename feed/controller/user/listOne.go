package user

import (
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/controller/util"
	"weibo/feed/service"
	"weibo/lib"
)

type ListOneReq struct{}

type ListOneResp struct {
	Nick  string `json:"nick"`
	Sex   byte   `json:"sex"`
	Email string `json:"email"`
}

// @prefilter("Auth")
// @postfilter("Boss")
func (user *User) ListOne(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.unsub_user"

	var listOneReq *ListOneReq

	//取用户id
	uid, err := util.GetUid(user)
	if err != nil {
		return
	}

	result, err := service.NewUser().GetById(uid)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, listOneReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	if result == nil {
		detail := "there is not this user"
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	resp := &ListOneResp{
		result.Nick,
		result.Sex,
		result.Email,
	}
	user.ReplyOk(w, resp)

	return
}
