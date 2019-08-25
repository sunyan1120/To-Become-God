package api

type Feed struct {
	ID        string `json:"id" bson:"_id"`
	UserId    string `json:"user_id"`
	Txt       string `json:"txt" bson:"_txt"`
	ViewTime  int    `json:"view_time" bson:"view_time"`
	CreatTime int64  `json:"creat_time"`
	UserNick  string `json:"user_nick"`
}

func NewFeed() *Feed {
	return &Feed{}
}
