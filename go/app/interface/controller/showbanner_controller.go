package controller

import (
	"html/template"
	"log"
	"net/http"

	"go-http-server/interface/repository"
	"go-http-server/usecase/interactor"
)

type showbannerController struct {
	Interactor interactor.ShowbannerInteractor
}

type ShowbannerController interface {
	View(http.ResponseWriter, *http.Request)
}

func NewShowbannerController(t *template.Template) ShowbannerController {
	return &showbannerController{
		Interactor: interactor.NewShowbannerInteractor(repository.NewShowbannerRepository(t)),
	}
}

func (sc *showbannerController) View(w http.ResponseWriter, r *http.Request) {
	if err := sc.Interactor.GetView(w, r); err != nil {
		log.Fatal(err)
	}
}
