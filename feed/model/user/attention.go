package user

import (
	rd "github.com/garyburd/redigo/redis"
	"github.com/globalsign/mgo"
	"weibo/feed/api"
	"weibo/feed/mongo"
	"weibo/feed/redis"
)

type Attention api.Attention

// Db 返回db name
func (attention *Attention) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (attention *Attention) Table() (table string) {
	return "attention"
}

// GetC 返回db col
func (attention *Attention) GetC() (c *mgo.Collection) {
	db, table := attention.Db(), attention.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
func (attention *Attention) GetConnection() (c rd.Conn) {
	r := redis.RDS["user"]
	c = r.Get()
	return
}
