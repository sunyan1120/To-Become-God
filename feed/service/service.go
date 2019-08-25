/*
Package service 用于定义服务层代码。
只允许在这里添加对外暴露的接口
*/
package service

import (
	"weibo/feed/service/comment"
	"weibo/feed/service/favor"
	"weibo/feed/service/feed"
	"weibo/feed/service/trends"
	"weibo/feed/service/user"
	"weibo/feed/service/view"
)

func NewFeed() *feed.Feed {
	return &feed.Feed{}
}
func NewUser() *user.User {
	return &user.User{}
}
func NewToken() *user.Token {
	return &user.Token{}
}
func NewAttention() *user.Attention {
	return &user.Attention{}
}
func NewFans() *user.Fans {
	return &user.Fans{}
}
func NewComment() *comment.Comment {
	return &comment.Comment{}
}
func NewView() *view.View {
	return &view.View{}
}
func NewFavor() *favor.Favor {
	return &favor.Favor{}
}
func NewTrends() *trends.Trends {
	return &trends.Trends{}
}
