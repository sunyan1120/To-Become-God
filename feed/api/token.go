package api

type Token struct {
	UserID  string `json:"userId"`
	Token   string `json:"token"`
	OutTime int    `json:"out_time"`
}

func NewToken() *Token {
	return &Token{}
}
