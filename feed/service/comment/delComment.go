package comment

import "weibo/feed/model"

func (comment *Comment) Del(id string) (err error) {
	commentModel := model.NewComment()
	commentModel.ID = id
	if err = commentModel.Del(); err != nil {
		return
	}

	return
}
