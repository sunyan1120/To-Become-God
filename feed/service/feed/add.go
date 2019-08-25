package feed

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

// Add 定义新增操作
func (feed *Feed) Add(feedApi *api.Feed) (err error) {
	feedModel := model.NewFeed()
	feedModel.UserId = feedApi.UserId
	feedModel.ID = feedApi.ID
	feedModel.Txt = feedApi.Txt
	feedModel.ViewTime = feedApi.ViewTime
	feedModel.CreatTime = feedApi.CreatTime
	feedModel.UserNick = feedApi.UserNick
	if err = feedModel.Add(); err != nil {
		return
	}

	fansModel := model.NewFans()
	fansModel.UserId = feedApi.UserId
	fansId, err := fansModel.GetFansId()

	for _, i := range fansId {
		feedModel.AddFansFeed(i)
	}

	return
}
