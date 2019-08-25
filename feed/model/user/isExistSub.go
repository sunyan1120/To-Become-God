package user

import "github.com/garyburd/redigo/redis"

func (attention *Attention) IsExistSub() (ret float64, err error) {
	c := attention.GetConnection()
	defer c.Close()

	ret, err = redis.Float64(c.Do("ZSCORE", attention.UserId+":s", attention.AttId))
	if err != nil {
		return
	}

	return

}
