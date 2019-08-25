package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (fans *Fans) GetAllFans(id string) (fansApi []*api.User, err error) {
	fansModel := model.NewFans()
	userModel := model.NewUser()
	fansModel.UserId = id

	var resp []*api.User

	fansId, err := fansModel.GetFansId()
	for _, i := range fansId {
		userModel.ID = i
		fansRet, _ := userModel.GetByIdApi()
		fansRet.Password = ""
		resp = append(resp, fansRet)
	}

	if fansModel == nil {
		return
	}

	fansApi = ([]*api.User)(resp)

	return
}

/*func (fans *Fans) GetFansCount(id string) (fansCount int, err error) {
	fansModel := model.NewFans()
	userModel := model.NewUser()
	fansModel.UserId = id

	var resp []*api.User

	fansId, err := fansModel.GetFansId()
	for _, i := range fansId {
		userModel.ID = i
		fansRet, _ := userModel.GetByIdApi()
		fansRet.Password = ""
		resp = append(resp, fansRet)
	}

	if fansModel == nil {
		return
	}

	fansApi = ([]*api.User)(resp)

	return
}*/
