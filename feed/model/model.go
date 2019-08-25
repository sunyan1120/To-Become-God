/*
Package model 用于模型层定义，所有db及cache对象封装均定义在这里。
只允许在这里添加对外暴露的接口
*/
package model

import (
	"weibo/feed/model/comment"
	"weibo/feed/model/favor"
	"weibo/feed/model/feed"
	"weibo/feed/model/trends"
	"weibo/feed/model/user"
	"weibo/feed/model/view"
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
