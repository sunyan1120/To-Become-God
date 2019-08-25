package user

// Add 定义新增操作
func (user *User) Register() (err error) {
	c := user.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(user)
	if err != nil {
		return
	}

	return
}
