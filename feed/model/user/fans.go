package user

import (
	rd "github.com/garyburd/redigo/redis"
	"github.com/globalsign/mgo"
	"weibo/feed/api"
	"weibo/feed/mongo"
	"weibo/feed/redis"
)

type Fans api.Fans

// Db 返回db name
func (fans *Fans) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (fans *Fans) Table() (table string) {
	return "fans"
}

// GetC 返回db col
func (fans *Fans) GetC() (c *mgo.Collection) {
	db, table := fans.Db(), fans.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
func (fans *Fans) GetConnection() (c rd.Conn) {
	r := redis.RDS["user"]
	c = r.Get()
	return
}
