package favor

import (
	"github.com/garyburd/redigo/redis"
	"weibo/feed/api"
	r "weibo/feed/redis"
)

type Favor api.Favor

func (favor *Favor) GetConnection() (c redis.Conn) {
	r := r.RDS["favor"]
	c = r.Get()
	return
}
