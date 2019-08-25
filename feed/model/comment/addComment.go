package comment

func (comment *Comment) AddComment() (err error) {
	c := comment.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(comment)
	if err != nil {
		return
	}

	return
}
