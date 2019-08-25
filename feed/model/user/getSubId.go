package user

import "github.com/garyburd/redigo/redis"

func (attention *Attention) GetSubId() (att []interface{}, err error) {
	c := attention.GetConnection()
	defer c.Close()

	att, err = redis.Values(c.Do("zrange", attention.UserId+":s", 0, -1))
	if err != nil {
		return
	}

	return
}
