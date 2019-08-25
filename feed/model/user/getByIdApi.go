package user

import (
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
	"weibo/feed/api"
)

func (user *User) GetByIdApi() (userRet *api.User, err error) {
	c := user.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"_id": user.ID}).One(&userRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
