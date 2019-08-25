package filter

import (
	"github.com/simplejia/clog"
	"net/http"
	"weibo/feed/service"
	"weibo/lib"
)

func Auth(w http.ResponseWriter, r *http.Request, m map[string]interface{}) bool {
	fun := "filter auth"

	//tokenCookie, err := r.Cookie("token")

	/*if err != nil {
		clog.Error("cookie no token")
		return false
	}*/

	//tokenStr := tokenCookie.Value
	tokenStr := r.Header.Get("token")
	uid, err := service.NewToken().GetTokens(tokenStr)
	if err != nil {
		clog.Error("%s param err: %v", fun, "token无效")
		w.Write([]byte("err:token 是无效的"))
		return false
	}
	if uid == "" {
		clog.Error("timeouts")
		return false
	}
	if controller, exist := m["__C__"]; !exist {
		clog.Error("controller has err")
		return false
	} else {
		uc := controller.(lib.IBase)
		uc.SetParam("uid", uid)
		return true
	}
}
