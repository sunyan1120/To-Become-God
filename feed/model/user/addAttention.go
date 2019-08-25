package user

import (
	"time"
)

func (attention *Attention) AddAttention() (err error) {
	c := attention.GetConnection()
	defer c.Close()

	_, err = c.Do("ZADD", attention.UserId+":s", time.Now().Unix(), attention.AttId)
	if err != nil {
		return
	}
	_, err = c.Do("ZADD", attention.AttId+":f", time.Now().Unix(), attention.UserId)
	if err != nil {
		return
	}

	return
}
