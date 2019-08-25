package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (user *User) Register(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.ID = userApi.ID
	userModel.Nick = userApi.Nick
	userModel.Email = userApi.Email
	userModel.Password = userApi.Password
	userModel.Sex = userApi.Sex
	if err = userModel.Register(); err != nil {
		return
	}

	return
}
