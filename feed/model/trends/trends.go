package trends

import (
	rd "github.com/garyburd/redigo/redis"
	mgo "github.com/globalsign/mgo"
	"weibo/feed/api"
	"weibo/feed/mongo"
	"weibo/feed/redis"
)

type Trends api.Trends

// Db 返回db name
func (trends *Trends) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (trends *Trends) Table() (table string) {
	return "feed"
}

// GetC 返回db col
func (trends *Trends) GetC() (c *mgo.Collection) {
	db, table := trends.Db(), trends.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
func (trends *Trends) GetConnection() (c rd.Conn) {
	r := redis.RDS["user"]
	c = r.Get()
	return
}
