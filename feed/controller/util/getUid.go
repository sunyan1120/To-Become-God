package util

import (
	"errors"
	"github.com/simplejia/clog"
	"weibo/lib"
)

func GetUid(user lib.IBase) (string, error) {
	var uid string
	uidParam, exists := user.GetParam("uid")
	if !exists {
		clog.Error("uid has err")
		return "", errors.New("uid has err")
	} else {
		uid = uidParam.(string)
		return uid, nil
	}
}
