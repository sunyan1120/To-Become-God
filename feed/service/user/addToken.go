package user

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (token *Token) AddToken(tokenApi *api.Token) (err error) {
	tokenModel := model.NewToken()
	tokenModel.UserID = tokenApi.UserID
	tokenModel.Token = tokenApi.Token
	if err = tokenModel.AddToken(); err != nil {
		return
	}

	return
}
