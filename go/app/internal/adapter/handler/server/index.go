package server

import (
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/usecase"
	"log"
	"net/http"
	"text/template"
)

type IndexHandler struct {
	getIndexPageData usecase.GetIndexPageData
}

func NewIndexHandler(getIndexPageData usecase.GetIndexPageData) *IndexHandler {
	return &IndexHandler{
		getIndexPageData: getIndexPageData,
	}
}

func (h *IndexHandler) Handle(responseWriter http.ResponseWriter, r *http.Request) {
	envValues := configuration.LoadEnv()
	title := envValues.IndexPage.Title
	header := envValues.IndexPage.Header

	output := h.getIndexPageData.Execute(usecase.GetIndexPageDataInput{
		Title:  title,
		Header: header,
	})
	if output.Error != nil {
		log.Fatal(output.Error)
	}

	t, err := template.ParseFiles(output.Template)
	if err != nil {
		log.Fatalf("Failed Create template. due to an error: %v\n", err)
	}

	if err := t.Execute(responseWriter, output.PageData); err != nil {
		log.Fatalf("Failed Execute template. due to an error: %v\n", err)
	}
}
