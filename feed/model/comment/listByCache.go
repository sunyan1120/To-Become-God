package comment

import (
	"github.com/garyburd/redigo/redis"
)

func (comment *Comment) ListByCahe() (comments []string, err error) {
	c := comment.GetConnection()
	defer c.Close()

	comments, err = redis.Strings(c.Do("zrevrange", comment.FeedId+":com", 0, -1))
	if err != nil {
		return
	}

	return
}
