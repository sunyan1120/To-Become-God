package favor

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (favor *Favor) Set(favorApi *api.Favor) (err error) {
	favorModel := model.NewFavor()
	favorModel.UserId = favorApi.UserId
	favorModel.FeedId = favorApi.FeedId
	favorModel.UserNick = favorApi.UserNick
	if err = favorModel.Set(); err != nil {
		return
	}

	return
}
