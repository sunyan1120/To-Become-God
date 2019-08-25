package user

import (
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/controller/util"
	"weibo/feed/service"
	"weibo/lib"
)

type GetFansReq struct{}

type GetFansResp struct {
	Fans []*api.User `json:"fans"`
}

// @prefilter("Auth")
// @postfilter("Boss")
func (user *User) GetFansUsers(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.Get_fans_users"

	//取用户id
	uid, err := util.GetUid(user)
	if err != nil {
		return
	}

	//count, err := service.NewFans()

	result, err := service.NewFans().GetAllFans(uid)
	if err != nil {
		clog.Error("%s fans.Get err: %v, req: %v", fun, err, nil)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &GetFansResp{
		result,
	}
	user.ReplyOk(w, resp)

	return
}
