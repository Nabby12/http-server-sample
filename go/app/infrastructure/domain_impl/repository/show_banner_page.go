package repository

import (
	"fmt"
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"html/template"
	"net/http"
)

type showBannerPageRepositoryImpl struct{}

func NewShowBannerPage() repository.ShowBannerPage {
	return &showBannerPageRepositoryImpl{}
}

const showBannerHtml string = "showbanner.html"

func (r *showBannerPageRepositoryImpl) SetView(showBannerPage model.ShowBannerPage, responseWriter http.ResponseWriter) error {
	envValues := configuration.LoadEnv()
	t, err := template.ParseFiles(fmt.Sprintf("%v/%v", envValues.PublicPath, showBannerHtml))
	if err != nil {
		fmt.Printf("Failed Create template. due to an error: %v\n", err)
		return err
	}

	if err := t.Execute(responseWriter, showBannerPage); err != nil {
		fmt.Printf("Failed Execute template. due to an error: %v\n", err)
		return err
	}

	return nil
}
