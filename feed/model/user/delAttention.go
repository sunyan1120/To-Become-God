package user

func (attention *Attention) DelAttention() (err error) {
	c := attention.GetConnection()
	defer c.Close()

	_, err = c.Do("ZREM", attention.UserId+":s", attention.AttId)
	if err != nil {
		return
	}
	_, err = c.Do("ZREM", attention.AttId+":f", attention.UserId)
	if err != nil {
		return
	}

	return
}
