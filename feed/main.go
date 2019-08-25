// skel，just for demo
// author: simplejia
// date: 2017/12/8

//go:generate wsp -s -d

package main

import (
	"fmt"
	"github.com/simplejia/clog"
	"github.com/simplejia/lc"
	"github.com/simplejia/utils"
	"html/template"
	"log"
	"net/http"
	"weibo/feed/conf"
	"weibo/lib"
)

func init() {
	lc.Init(1e5)

	clog.AddrFunc = func() (string, error) {
		return lib.NameWrap(conf.C.Addrs.Clog)
	}
	clog.Init(conf.C.Clog.Name, "", conf.C.Clog.Level, conf.C.Clog.Mode)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clog.Error("%s is not found", r.RequestURI)
		http.NotFound(w, r)
	})
}

func main() {
	fun := "main"
	clog.Info(fun)

	addr := fmt.Sprintf("%s:%d", "0.0.0.0", conf.C.App.Port)
	http.HandleFunc("/index.html", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("./web/index.html")
		if err != nil {
			log.Println("err:", err)
			return
		}
		t.Execute(res, nil)
	})
	err := utils.ListenAndServe(addr, nil)
	if err != nil {
		clog.Error("%s err: %v, addr: %v", fun, err, addr)
	}
}
