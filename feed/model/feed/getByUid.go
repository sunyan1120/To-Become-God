package feed

import (
	"fmt"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
	"weibo/feed/api"
)

func (feed *Feed) GetByUid() (feedRet []*api.Feed, err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	fmt.Println(feed.UserId, "uid为！")
	err = c.Find(bson.M{"userid": feed.UserId}). /*Sort("creattime", "-1").Skip(0).Limit(20).*/ All(&feedRet)

	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
