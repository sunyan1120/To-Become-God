package api

type Trends struct {
	FeedId string `json:"feed_id"`
	Txt    string `json:"txt"`
	Ct     int64  `json:"ct"`
	Uv     int    `json:"uv"`
	Favor  struct {
		Count   int `json:"count"`
		IsFavor int `json:"is_favor"`
	} `json:"favor"`
	User struct {
		Id   string `json:"id"`
		Nick string `json:"nick"`
	} `json:"user"`
}

func NewTrends() *Trends {
	return &Trends{}
}
