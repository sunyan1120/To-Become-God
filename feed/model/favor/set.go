package favor

import "time"

func (favor *Favor) Set() (err error) {
	c := favor.GetConnection()
	defer c.Close()

	_, err = c.Do("ZADD", favor.FeedId+":fav",
		time.Now().Unix(), favor.UserId+":"+favor.UserNick)
	if err != nil {
		return
	}

	return
}
