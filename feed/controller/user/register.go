package user

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

// RegisterReq 定义输入
type RegisterReq struct {
	Nick     string `json:"nick"`
	Sex      byte   `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Regular 用于参数校验
func (registerReq *RegisterReq) Regular() (ok bool) {
	if registerReq == nil {
		return
	}

	if "" == registerReq.Nick || "" == registerReq.Email || "" == registerReq.Password ||
		(registerReq.Sex != 1 && registerReq.Sex != 2) {
		return
	}

	ok = true
	return
}

// RegisterResp 定义输出
type RegisterResp struct{}

// @postfilter("Boss")
func (user *User) Register(w http.ResponseWriter, r *http.Request) {
	fun := "feed.User.Rigister"

	var registerReq *RegisterReq
	if err := json.Unmarshal(user.ReadBody(r), &registerReq); err != nil || !registerReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, registerReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}
	userApi := api.NewUser()
	//判断该email是否已经注册过
	results, err := service.NewUser().GetOne(registerReq.Email)
	if err != nil {
		clog.Error("%s user.register err: %v, req: %v", fun, err, registerReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	if results != nil {
		detail := "User already exists"
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}
	id := uuid.Must(uuid.NewV4())
	userApi.ID = id.String()
	userApi.Nick = registerReq.Nick
	userApi.Email = registerReq.Email
	userApi.Password = util.Encryption(registerReq.Password) //使用md5加密
	userApi.Sex = registerReq.Sex

	if err := service.NewUser().Register(userApi); err != nil {
		clog.Error("%s user.register err: %v, req: %v", fun, err, registerReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &RegisterResp{}
	user.ReplyOk(w, resp)

	return
}
