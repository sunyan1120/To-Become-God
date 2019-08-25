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

type UpdateReq struct {
	Nick     string `json:"nick"`
	Sex      byte   `json:"sex"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil {
		return
	}

	if updateReq.Sex != 0 && updateReq.Sex != 1 && updateReq.Sex != 2 {
		return
	}

	ok = true
	return
}

type UpdateResp struct{}

// @prefilter("Auth")
// @postfilter("Boss")
func (user *User) Update(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.Update"

	var updateReq *UpdateReq
	if err := json.Unmarshal(user.ReadBody(r), &updateReq); err != nil || !updateReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}
	//取用户id
	uid, err := util.GetUid(user)
	if err != nil {
		return
	}
	results, err := service.NewUser().GetById(uid)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	if results == nil {
		detail := "there is not this user"
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}
	userApi := api.NewUser()
	userApi.ID = uid
	userApi.Sex = updateReq.Sex
	userApi.Nick = updateReq.Nick
	if updateReq.Password != "" {
		userApi.Password = util.Encryption(updateReq.Password)
	}
	userApi.Email = updateReq.Email
	if err := service.NewUser().Update(userApi); err != nil {
		clog.Error("%s feed.Update err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &UpdateResp{}
	user.ReplyOk(w, resp)

	return
}
