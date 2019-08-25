package feed

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
	"weibo/feed/controller/util"

	"weibo/feed/api"
	"weibo/feed/service"
	"weibo/lib"

	"github.com/simplejia/clog"
)

// AddReq 定义输入
type AddReq struct {
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	if "" == addReq.Txt {
		return
	}

	ok = true
	return
}

// AddResp 定义输出
type AddResp struct {
	data map[string]interface{}
}

// @prefilter("Auth")
// @postfilter("Boss")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Add"

	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}

	//取用户id
	uid, err := util.GetUid(feed)
	if err != nil {
		return
	}
	//根据id去用户
	user, err := service.NewUser().GetById(uid)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}
	if user == nil {
		detail := "there is not this user"
		feed.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	feedApi := api.NewFeed()
	id := uuid.Must(uuid.NewV4())
	feedApi.ID = id.String()
	feedApi.UserId = uid
	feedApi.Txt = addReq.Txt
	feedApi.ViewTime = 0
	feedApi.CreatTime = time.Now().Unix()
	feedApi.UserNick = user.Nick
	if err := service.NewFeed().Add(feedApi); err != nil {
		clog.Error("%s skel.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddResp{map[string]interface{}{"data": feedApi}}
	feed.ReplyOk(w, resp)

	return
}
