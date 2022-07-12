package controller

import (
	"html/template"
	"log"
	"net/http"

	"go-http-server/interface/repository"
	"go-http-server/usecase/interactor"
)

type indexController struct {
	Interactor interactor.IndexInteractor
}

type IndexController interface {
	View(http.ResponseWriter, *http.Request)
}

func NewIndexController(t *template.Template) IndexController {
	return &indexController{
		Interactor: interactor.NewIndexInteractor(repository.NewIndexRepository(t)),
	}
}

func (ic *indexController) View(w http.ResponseWriter, r *http.Request) {
	if err := ic.Interactor.GetView(w, r); err != nil {
		log.Fatal(err)
	}
}
