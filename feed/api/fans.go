package api

type Fans struct {
	UserId string `json:"user_id" bson:"user_id"`
	FansId string `json:"fans_id" bson:"fans_id"`
}

func NewFans() *Fans {
	return &Fans{}
}
