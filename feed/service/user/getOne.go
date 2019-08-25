package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (user *User) GetOne(email string) (userApi *api.User, err error) {
	userModel := model.NewUser()
	userModel.Email = email
	if userModel, err = userModel.GetOne(); err != nil {
		return
	}
	if userModel == nil {
		return
	}
	userApi = (*api.User)(userModel)
	return
}
