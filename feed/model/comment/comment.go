package comment

import (
	"github.com/garyburd/redigo/redis"
	"github.com/globalsign/mgo"
	"weibo/feed/api"
	"weibo/feed/mongo"
	r "weibo/feed/redis"
)

type Comment api.Comment

// Db 返回db name
func (comment *Comment) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (comment *Comment) Table() (table string) {
	return "comment"
}

// GetC 返回db col
func (comment *Comment) GetC() (c *mgo.Collection) {
	db, table := comment.Db(), comment.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
func (comment *Comment) GetConnection() (c redis.Conn) {
	r := r.RDS["favor"]
	c = r.Get()
	return
}
