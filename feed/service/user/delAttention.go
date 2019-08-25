package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (attention *Attention) DelAttention(attentionApi *api.Attention) (err error) {
	attentionModel := model.NewAttention()
	attentionModel.AttId = attentionApi.AttId
	attentionModel.UserId = attentionApi.UserId
	if err = attentionModel.DelAttention(); err != nil {
		return
	}

	return
}
