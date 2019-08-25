package trends

import (
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/api"
	"weibo/feed/controller/util"
	"weibo/feed/service"
	"weibo/lib"
)

type GetTrendsReq struct{}

type GetTrendsResp struct {
	Trends []*api.Trends `json:"attentions"`
}

// @prefilter("Auth")]
// @postfilter("Boss")
func (trends *Trends) GetUserTrends(w http.ResponseWriter, r *http.Request) {
	fun := "trends.Trends.GetUserTrends"

	//取用户id
	uid, err := util.GetUid(trends)
	if err != nil {
		return
	}

	result, err := service.NewTrends().GetUserTrends(uid)
	if err != nil {
		clog.Error("%s Attention.Get err: %v, req: %v", fun, err, nil)
		trends.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &GetTrendsResp{
		result,
	}
	trends.ReplyOk(w, resp)

	return
}
