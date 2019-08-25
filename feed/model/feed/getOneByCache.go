package feed

import (
	"github.com/garyburd/redigo/redis"
)

func (feed *Feed) GetOneByCahe() (feedstr string, err error) {
	c := feed.GetConnection()
	defer c.Close()

	feedstr, err = redis.String(c.Do("get", feed.ID+":feed"))
	if err != nil {
		return
	}

	return
}
