package user

import (
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/controller/util"
	"weibo/feed/service"
	"weibo/lib"
)

type GetAttentionsReq struct{}

type GetAttentionsResp struct {
	Attentions []*api.User `json:"attentions"`
}

// @prefilter("Auth")
// @postfilter("Boss")
func (user *User) GetSubUsers(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.Get_sub_users"

	//取用户id
	uid, err := util.GetUid(user)
	if err != nil {
		return
	}

	result, err := service.NewAttention().GetAllSubUsers(uid)
	if err != nil {
		clog.Error("%s Attention.Get err: %v, req: %v", fun, err, nil)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &GetAttentionsResp{
		result,
	}
	user.ReplyOk(w, resp)

	return
}
