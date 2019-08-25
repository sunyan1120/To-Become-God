package comment

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (comment *Comment) AddComment(commentApi *api.Comment) (err error) {
	commentModel := model.NewComment()
	commentModel.ID = commentApi.ID
	commentModel.PostUserId = commentApi.PostUserId
	commentModel.FeedId = commentApi.FeedId
	commentModel.CommentTxt = commentApi.CommentTxt
	commentModel.UserNick = commentApi.UserNick
	if err = commentModel.AddComment(); err != nil {
		return
	}

	err = commentModel.AddCache(commentApi)
	if err != nil {
		return
	}

	return
}
