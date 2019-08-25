package api

type User struct {
	ID       string `json:"id" bson:"_id"`
	Nick     string `json:"nick" bson:"nick"`
	Sex      byte   `json:"sex" bson:"sex"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func NewUser() *User {
	return &User{}
}
