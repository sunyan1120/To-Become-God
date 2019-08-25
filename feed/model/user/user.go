package user

import (
	mgo "github.com/globalsign/mgo"
	"weibo/feed/api"
	"weibo/feed/mongo"
)

type User api.User

// Db 返回db name
func (user *User) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (user *User) Table() (table string) {
	return "user"
}

// GetC 返回db col
func (user *User) GetC() (c *mgo.Collection) {
	db, table := user.Db(), user.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
