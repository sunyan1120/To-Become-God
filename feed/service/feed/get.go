package feed

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

// Get 定义获取操作
func (feed *Feed) Get(id string) (feedApi *api.Feed, err error) {
	feedModel := model.NewFeed()
	feedModel.ID = id

	if feedModel, err = feedModel.Get(); err != nil {
		return
	}
	if feedModel == nil {
		return
	}

	feedApi = (*api.Feed)(feedModel)

	return
}
