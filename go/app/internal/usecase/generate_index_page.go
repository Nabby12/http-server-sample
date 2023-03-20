package usecase

import (
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"net/http"
)

type (
	generateIndexPageImpl struct {
		indexPageRepository repository.IndexPage
	}
	GenerateIndexPage interface {
		Execute(input GenerateIndexPageInput) *GenerateIndexPageOutput
	}
	GenerateIndexPageInput struct {
		IndexPage      model.IndexPage
		ResponseWriter http.ResponseWriter
	}
	GenerateIndexPageOutput struct {
		Error error
	}
)

func NewGenerateIndexPage(indexPageRepository repository.IndexPage) GenerateIndexPage {
	return &generateIndexPageImpl{
		indexPageRepository: indexPageRepository,
	}
}

func (u *generateIndexPageImpl) Execute(input GenerateIndexPageInput) *GenerateIndexPageOutput {
	if err := u.indexPageRepository.SetView(input.IndexPage, input.ResponseWriter); err != nil {
		return &GenerateIndexPageOutput{
			Error: err,
		}
	}

	return &GenerateIndexPageOutput{
		Error: nil,
	}
}
