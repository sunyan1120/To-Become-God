package api

type Attention struct {
	UserId string `json:"user_id" bson:"user_id"`
	AttId  string `json:"att_id" bson:"att_id"`
}

func NewAttention() *Attention {
	return &Attention{}
}
