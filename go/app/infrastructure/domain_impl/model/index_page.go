package model

import "go-http-server/internal/domain/model"

type indexPage struct {
	title  string
	header string
}

func NewIndexPage(title string, header string) model.IndexPage {
	return &indexPage{
		title:  title,
		header: header,
	}
}

func (ip *indexPage) Title() string {
	return ip.title
}

func (ip *indexPage) Header() string {
	return ip.header
}
