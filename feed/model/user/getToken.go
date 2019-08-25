package user

import (
	"github.com/garyburd/redigo/redis"
)

func (token *Token) GetTokens() (uid string, err error) {
	c := token.GetConnection()
	defer c.Close()

	uid, err = redis.String(c.Do("GET", token.Token))
	if err == nil {
		return
	}
	return
}
