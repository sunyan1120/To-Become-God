package feed

import (
	"encoding/json"
	"weibo/feed/api"
)

func (feed *Feed) AddCache(thisFeed *api.Feed) (err error) {
	c := feed.GetConnection()
	defer c.Close()

	bt, err := json.Marshal(thisFeed)

	_, err = c.Do("set", feed.ID+":feed", string(bt))
	if err != nil {
		return
	}

	return
}
