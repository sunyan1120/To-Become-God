package favor

func (favor *Favor) UnSet() (err error) {
	c := favor.GetConnection()
	defer c.Close()

	_, err = c.Do("ZREM", favor.FeedId+":fav", favor.UserId+":"+favor.UserNick)
	if err != nil {
		return
	}

	return
}
