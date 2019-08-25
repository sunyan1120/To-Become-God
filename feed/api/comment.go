package api

type Comment struct {
	ID         string `json:"id" bson:"_id"`
	FeedId     string `json:"feed_id"`
	PostUserId string `json:"post_user_id"`
	CommentTxt string `json:"comment_txt"`
	UserNick   string `json:"user_nick"`
}

func NewComment() *Comment {
	return &Comment{}
}
