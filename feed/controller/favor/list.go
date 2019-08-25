package favor

import (
	"encoding/json"
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/service"
	"weibo/lib"
)

type ListReq struct {
	Id     string `json:"id"`
	OffSet int    `json:"off_set"`
	Limit  int    `json:"limit"`
}

func (listReq *ListReq) Regular() (ok bool) {
	if listReq == nil {
		return
	}

	if "" == listReq.Id {
		return
	}

	ok = true
	return
}

type ListResp struct {
	//Favors []*api.User `json:"favors"`
	UserId   string `json:"user_id"`
	UserNick string `json:"user_nick"`
}

// @prefilt("Auth")
// @postfilter("Boss")
func (favor *Favor) List(w http.ResponseWriter, r *http.Request) {
	fun := "favor.Favor.List"

	var listReq *ListReq
	//获得body体里面的参数，并校验
	if err := json.Unmarshal(favor.ReadBody(r), &listReq); err != nil || !listReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, listReq)
		favor.ReplyFail(w, lib.CodePara)
		return
	}

	result, err := service.NewFavor().List(listReq.Id, listReq.OffSet, listReq.Limit)
	if err != nil {
		clog.Error("%s Attention.Get err: %v, req: %v", fun, err, nil)
		favor.ReplyFail(w, lib.CodeSrv)
		return
	}

	ListResps := make([]ListResp, len(result))
	for i := 0; i < len(result); i++ {
		ListResps[i].UserId = result[i].ID
		ListResps[i].UserNick = result[i].Nick
	}

	favor.ReplyOk(w, &ListResps)

	return
}
