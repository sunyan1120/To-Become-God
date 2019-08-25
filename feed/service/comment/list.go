package comment

import (
	"encoding/json"
	"weibo/feed/api"
	"weibo/feed/model"
)

func (comment *Comment) GetAll(feedId string) (commentApi []*api.Comment, err error) {
	commentModel := model.NewComment()
	commentModel.FeedId = feedId
	var resp []*api.Comment

	//从缓存中取
	coms, err := commentModel.ListByCahe()
	//如果有直接返回
	if coms != nil {
		for _, v := range coms {
			com := api.NewComment()
			err = json.Unmarshal([]byte(v), com)
			if err != nil {
				return
			}
			resp = append(resp, com)
		}
		commentApi = ([]*api.Comment)(resp)
		return
	}
	//如果缓存中没有，去数据库中取
	if resp, err = commentModel.GetAll(); err != nil {
		return
	}
	//加入到缓存中
	for _, v := range resp {
		commentModel.FeedId = v.FeedId
		err = commentModel.AddCache(v)
		if err != nil {
			return
		}
	}

	commentApi = ([]*api.Comment)(resp)

	return
}
