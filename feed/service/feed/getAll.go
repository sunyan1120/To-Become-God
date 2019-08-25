package feed

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

// Get 定义获取操作
func (feed *Feed) GetAll() (feedApi []*api.Feed, err error) {
	feedModel := model.NewFeed()
	var resp []*api.Feed
	if resp, err = feedModel.GetAll(); err != nil {
		return
	}

	if feedModel == nil {
		return
	}

	feedApi = ([]*api.Feed)(resp)

	return
}
