package user

import (
	"weibo/feed/model"
)

func (attention *Attention) IsExistSub(id string, attId string) (b float64, err error) {
	attentionModel := model.NewAttention()
	attentionModel.UserId = id
	attentionModel.AttId = attId
	if b, err = attentionModel.IsExistSub(); err != nil {
		return
	}

	return
}
