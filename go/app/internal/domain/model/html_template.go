package model

type HtmlTemplate interface {
	Name() string
	Path() string
	FullPath() string
}
