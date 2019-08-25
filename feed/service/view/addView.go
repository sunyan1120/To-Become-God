package view

import (
	"weibo/feed/api"
	"weibo/feed/model"
)

func (view *View) Add(viewApi *api.Feed) (err error) {
	viewModel := model.NewView()
	viewModel.ID = viewApi.ID
	viewModel.ViewTime = viewApi.ViewTime
	if err = viewModel.Add(); err != nil {
		return
	}

	return
}
