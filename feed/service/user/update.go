package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (user *User) Update(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.ID = userApi.ID
	userModel.Password = userApi.Password
	userModel.Sex = userApi.Sex
	userModel.Nick = userApi.Nick
	userModel.Email = userApi.Email

	if err = userModel.Update(); err != nil {
		return
	}

	return
}
