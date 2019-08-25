package comment

import (
	"encoding/json"
	"time"
	"weibo/feed/api"
)

func (comment *Comment) AddCache(thisCom *api.Comment) (err error) {
	c := comment.GetConnection()
	defer c.Close()

	bt, err := json.Marshal(thisCom)

	_, err = c.Do("ZADD", comment.FeedId+":com", time.Now().Unix(), string(bt))
	if err != nil {
		return
	}

	return
}
func (comment *Comment) DelCache(thisCom *api.Comment) (err error) {
	c := comment.GetConnection()
	defer c.Close()

	bt, err := json.Marshal(thisCom)
	_, err = c.Do("ZREM", comment.FeedId+":com", string(bt))
	if err != nil {
		return
	}

	return
}
