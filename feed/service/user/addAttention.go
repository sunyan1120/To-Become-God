package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (attention *Attention) AddAttention(attentionApi *api.Attention) (err error) {
	attentionModel := model.NewAttention()
	attentionModel.UserId = attentionApi.UserId
	attentionModel.AttId = attentionApi.AttId
	//添加关注
	if err = attentionModel.AddAttention(); err != nil {
		return
	}

	//过的关注人的动态
	feedModel := model.NewFeed()
	feedModel.UserId = attentionApi.AttId
	feeds, err := feedModel.GetByUid()

	if err != nil {
		return
	}
	if feeds == nil {
		return
	}
	//加到缓存中
	for _, v := range feeds {
		feedModel.ID = v.ID
		err = feedModel.AddFansFeed(attentionApi.UserId)
		if err != nil {
			return
		}
	}

	return
}
