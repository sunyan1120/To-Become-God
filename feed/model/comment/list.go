package comment

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"weibo/feed/api"
)

func (comment *Comment) GetAll() (commentRet []*api.Comment, err error) {
	c := comment.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"feedid": comment.FeedId}).All(&commentRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
