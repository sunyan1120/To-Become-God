package view

import (
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

func (view *View) Add() (err error) {
	c := view.GetC()
	defer c.Database.Session.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"view_time": 1}},
		ReturnNew: true,
	}
	_, err = c.Find(bson.M{"_id": view.ID}).Apply(change, &view)

	if err != nil {
		return
	}

	return
}
