package interactor

import (
	"net/http"

	"go-http-server/usecase/repository"
)

type indexInteractor struct {
	Repository repository.IndexRepository
}

type IndexInteractor interface {
	GetView(http.ResponseWriter, *http.Request) error
}

func NewIndexInteractor(ir repository.IndexRepository) IndexInteractor {
	return &indexInteractor{
		Repository: ir,
	}
}

func (ii *indexInteractor) GetView(w http.ResponseWriter, r *http.Request) error {
	if err := ii.Repository.SetView(w, r); err != nil {
		return err
	}
	return nil
}
