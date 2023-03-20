package repository

import (
	"fmt"
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"html/template"
	"net/http"
)

type indexPageImpl struct{}

func NewIndexPage() repository.IndexPage {
	return &indexPageImpl{}
}

const indexHtml string = "index.html"

func (r *indexPageImpl) SetView(indexPage model.IndexPage, responseWriter http.ResponseWriter) error {
	envValues := configuration.LoadEnv()
	t, err := template.ParseFiles(fmt.Sprintf("%v/%v", envValues.PublicPath, indexHtml))
	if err != nil {
		fmt.Printf("Failed Create template. due to an error: %v\n", err)
		return err
	}

	if err := t.Execute(responseWriter, indexPage); err != nil {
		fmt.Printf("Failed Execute template. due to an error: %v\n", err)
		return err
	}

	return nil
}
