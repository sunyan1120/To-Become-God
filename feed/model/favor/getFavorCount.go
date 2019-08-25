package favor

import "github.com/garyburd/redigo/redis"

func (favor *Favor) GetFavorCount() (count int, err error) {
	c := favor.GetConnection()
	defer c.Close()

	count, err = redis.Int(c.Do("zcard", favor.FeedId+":fav"))

	if err != nil {
		return
	}

	return
}
