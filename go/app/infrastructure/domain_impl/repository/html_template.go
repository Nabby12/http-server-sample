package repository

import (
	"errors"
	inframodel "go-http-server/infrastructure/domain_impl/model"
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
)

type htmlTemplateRepositoryImpl struct{}

func NewHtmlTemplate() repository.HtmlTemplate {
	return &htmlTemplateRepositoryImpl{}
}

var (
	IndexKey      = "index"
	ShowBannerKey = "showbanner"
)

func (r *htmlTemplateRepositoryImpl) GetByName(key string) (model.HtmlTemplate, error) {
	envValues := configuration.LoadEnv()

	var name string
	var path string
	switch key {
	case IndexKey:
		name = "index.html"
		path = envValues.PublicPath
	case ShowBannerKey:
		name = "showbanner.html"
		path = envValues.PublicPath
	default:
		return nil, errors.New("failed to match key")
	}

	htmlTemplate := inframodel.NewHtmlTemplate(name, path)
	return htmlTemplate, nil
}
