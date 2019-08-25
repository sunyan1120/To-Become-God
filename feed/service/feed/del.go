package feed

import (
	"weibo/feed/model"
)

// Del 定义删除操作
func (feed *Feed) Del(id string) (err error) {
	feedModel := model.NewFeed()
	feedModel.ID = id
	if err = feedModel.Del(); err != nil {
		return
	}

	return
}
