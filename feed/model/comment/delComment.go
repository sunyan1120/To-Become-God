package comment

import (
	"github.com/globalsign/mgo"
)

func (comment *Comment) Del() (err error) {
	c := comment.GetC()
	defer c.Database.Session.Close()

	err = c.RemoveId(comment.ID)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
