package user

import "github.com/garyburd/redigo/redis"

func (fans *Fans) GetFansId() (att []string, err error) {
	c := fans.GetConnection()
	defer c.Close()

	att, err = redis.Strings(c.Do("zrange", fans.UserId+":f", 0, -1))
	if err != nil {
		return
	}

	return
}
