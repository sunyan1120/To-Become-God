package api

type Favor struct {
	FeedId   string `json:"article_id"`
	UserId   string `json:"user_id"`
	Count    int    `json:"count"`
	UserNick string `json:"user_nick"`
}

func NewFavor() *Favor {
	return &Favor{}
}
