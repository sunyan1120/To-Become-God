package favor

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (favor *Favor) UnSet(favorApi *api.Favor) (err error) {
	favorModel := model.NewFavor()
	favorModel.UserId = favorApi.UserId
	favorModel.FeedId = favorApi.FeedId
	favorModel.UserNick = favorApi.UserNick
	if err = favorModel.UnSet(); err != nil {
		return
	}

	return
}
