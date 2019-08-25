package trends

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (trends *Trends) GetUserTrends(id string) (trendsApi []*api.Trends, err error) {
	feedModel := model.NewFeed()
	favorModel := model.NewFavor()

	var resp = []*api.Trends{}

	trendsModel := model.NewTrends()
	//过去关注人动态
	result, err := trendsModel.GetUserTrends(id)
	if err != nil {
		return
	}
	if result == nil {
		return
	}
	for _, v := range result {
		trModel := api.NewTrends()

		feedModel.ID = v
		feed, _ := feedModel.Get()
		trModel.FeedId = v
		trModel.Txt = feed.Txt
		trModel.Ct = feed.CreatTime
		trModel.Uv = feed.ViewTime
		trModel.User.Id = feed.UserId
		trModel.User.Nick = feed.UserNick

		favorModel.FeedId = v
		count, _ := favorModel.GetFavorCount()
		trModel.Favor.Count = count
		if count != 0 {
			trModel.Favor.IsFavor = 1
		}

		resp = append(resp, trModel)
	}

	if resp == nil {
		return
	}
	trendsApi = ([]*api.Trends)(resp)

	return
}
