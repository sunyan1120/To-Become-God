package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (user *User) GetById(id string) (userApi *api.User, err error) {
	userModel := model.NewUser()
	userModel.ID = id
	if userModel, err = userModel.GetById(); err != nil {
		return
	}
	if userModel == nil {
		return
	}
	userApi = (*api.User)(userModel)
	return
}
