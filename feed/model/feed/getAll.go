package feed

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"weibo/feed/api"
)

// Get 定义获取操作
func (feed *Feed) GetAll() (feedRet []*api.Feed, err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{}).All(&feedRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
