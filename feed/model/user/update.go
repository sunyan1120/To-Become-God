package user

import (
	"gopkg.in/mgo.v2/bson"
)

func (user *User) Update() (err error) {
	c := user.GetC()
	defer c.Database.Session.Close()

	q := bson.M{}
	if user.Nick != "" {
		q["nick"] = user.Nick
	}
	if user.Sex == 1 || user.Sex == 2 {
		q["sex"] = user.Sex
	}
	if user.Password != "" {
		q["password"] = user.Password
	}
	if user.Email != "" {
		q["email"] = user.Email
	}

	err = c.UpdateId(user.ID, bson.M{"$set": q})
	if err != nil {
		return
	}

	return
}
