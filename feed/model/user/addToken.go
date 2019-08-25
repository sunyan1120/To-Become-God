package user

func (token *Token) AddToken() (err error) {
	c := token.GetConnection()
	defer c.Close()

	_, err = c.Do("SET", token.Token, token.UserID, "EX", 12000)
	if err != nil {
		return
	}

	return
}
