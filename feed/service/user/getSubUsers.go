package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (attention *Attention) GetAllSubUsers(id string) (attentionApi []*api.User, err error) {
	attentionModel := model.NewAttention()
	userModel := model.NewUser()
	attentionModel.UserId = id

	var resp []*api.User

	subs, err := attentionModel.GetSubId()
	for _, i := range subs {
		userModel.ID = string(i.([]uint8))
		attentionRet, _ := userModel.GetByIdApi()
		attentionRet.Password = ""
		resp = append(resp, attentionRet)
	}

	if attentionModel == nil {
		return
	}

	attentionApi = ([]*api.User)(resp)

	return
}
