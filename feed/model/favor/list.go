package favor

import (
	"github.com/garyburd/redigo/redis"
)

func (favor *Favor) List(offset int, limit int) (favors []string, err error) {
	c := favor.GetConnection()
	defer c.Close()

	favors, err = redis.Strings(c.Do("zrevrange", favor.FeedId+":fav", offset, limit))

	if err != nil {
		return
	}

	return
}
