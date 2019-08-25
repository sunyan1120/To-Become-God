package user

import (
	"weibo/feed/model"
)

func (token *Token) GetTokens(tok string) (uid string, err error) {
	tokenModel := model.NewToken()
	tokenModel.Token = tok

	if uid, err = tokenModel.GetTokens(); err != nil {
		return
	}
	return
}
