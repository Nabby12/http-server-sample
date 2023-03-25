package model

import (
	"fmt"
	"go-http-server/internal/domain/model"
)

type htmlTemplate struct {
	name string
	path string
}

func NewHtmlTemplate(name string, path string) model.HtmlTemplate {
	return &htmlTemplate{
		name: name,
		path: path,
	}
}

func (ht *htmlTemplate) Name() string {
	return ht.name
}

func (ht *htmlTemplate) Path() string {
	return ht.path
}

func (ht *htmlTemplate) FullPath() string {
	return fmt.Sprintf("%v/%v", ht.path, ht.name)
}
