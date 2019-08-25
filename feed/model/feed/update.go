package feed

import "gopkg.in/mgo.v2/bson"

func (feed *Feed) Update() (err error) {
	c := feed.GetC()
	defer c.Database.Session.Close()

	txt := bson.M{"$set": bson.M{"_txt": feed.Txt}}
	err = c.UpdateId(feed.ID, txt)
	if err != nil {
		return
	}

	return
}
