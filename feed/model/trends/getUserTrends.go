package trends

import "github.com/garyburd/redigo/redis"

func (trends *Trends) GetUserTrends(id string) (trs []string, err error) {
	c := trends.GetConnection()
	defer c.Close()

	trs, err = redis.Strings(c.Do("zrange", id+":SubFeed", 0, -1))
	if err != nil {
		return
	}

	return
}
