package server

import (
	"go-http-server/infrastructure/domain_impl/model"
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/usecase"
	"log"
	"net/http"
)

type IndexHandler struct {
	generateIndexPage usecase.GenerateIndexPage
}

func NewIndexHandler(generateIndexPage usecase.GenerateIndexPage) *IndexHandler {
	return &IndexHandler{
		generateIndexPage: generateIndexPage,
	}
}

func (h *IndexHandler) Handle(responseWriter http.ResponseWriter, r *http.Request) {
	envValues := configuration.LoadEnv()
	title := envValues.IndexPage.Title
	header := envValues.IndexPage.Header

	indexPage := model.NewIndexPage(title, header)

	if output := h.generateIndexPage.Execute(usecase.GenerateIndexPageInput{
		IndexPage:      indexPage,
		ResponseWriter: responseWriter,
	}); output.Error != nil {
		log.Fatal(output.Error)
	}
}
