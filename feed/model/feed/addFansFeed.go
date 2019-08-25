package feed

import "time"

func (feed *Feed) AddFansFeed(i string) (err error) {
	c := feed.GetConnection()
	defer c.Close()

	_, err = c.Do("ZADD", i+":SubFeed", time.Now().Unix(), feed.ID)
	if err != nil {
		return
	}

	return
}
