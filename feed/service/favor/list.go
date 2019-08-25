package favor

import (
	"strings"
	"weibo/feed/api"
	"weibo/feed/model"
)

func (favor *Favor) List(feedId string, offset int, limit int) (favorApi []*api.User, err error) {
	favorModel := model.NewFavor()
	favorModel.FeedId = feedId
	var resp []*api.User

	//拿到点赞人的id和nick的字符串切片
	favors, err := favorModel.List(offset, limit)

	respChan := make(chan bool, len(favors))
	for _, i := range favors {
		go func(i string) {
			userModel := api.NewUser()
			//将每个字符串以：切割
			users := strings.Split(i, ":")
			userModel.ID = users[0]
			userModel.Nick = users[1]
			resp = append(resp, userModel)

			respChan <- true
		}(i)

	}
	for i := 0; i < len(favors); i++ {
		<-respChan
	}

	if favorModel == nil {
		return
	}

	favorApi = ([]*api.User)(resp)

	return
}
