package view

import (
	mgo "github.com/globalsign/mgo"
	"weibo/feed/api"
	"weibo/feed/mongo"
)

type View api.Feed

// Db 返回db name
func (view *View) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (view *View) Table() (table string) {
	return "feed"
}

// GetC 返回db col
func (view *View) GetC() (c *mgo.Collection) {
	db, table := view.Db(), view.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
