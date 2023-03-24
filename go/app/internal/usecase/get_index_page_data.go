package usecase

import (
	inframodel "go-http-server/infrastructure/domain_impl/model"
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
)

type (
	getIndexPageDataImpl struct {
		HtmlTemplateRepository repository.HtmlTemplate
	}
	GetIndexPageData interface {
		Execute(input GetIndexPageDataInput) *GetIndexPageDataOutput
	}
	GetIndexPageDataInput struct {
		Title  string
		Header string
	}
	GetIndexPageDataOutput struct {
		PageData model.IndexPage
		Template string
		Error    error
	}
)

var (
	IndexKey = "index"
)

func NewGetIndexPageData(HtmlTemplateRepository repository.HtmlTemplate) GetIndexPageData {
	return &getIndexPageDataImpl{
		HtmlTemplateRepository: HtmlTemplateRepository,
	}
}

func (u *getIndexPageDataImpl) Execute(input GetIndexPageDataInput) *GetIndexPageDataOutput {
	indexPage := inframodel.NewIndexPage(input.Title, input.Header)

	htmlTemplate, err := u.HtmlTemplateRepository.GetByName(IndexKey)
	if err != nil {
		return &GetIndexPageDataOutput{
			PageData: nil,
			Template: "",
			Error:    err,
		}
	}
	htmlTemplateFullPath := htmlTemplate.FullPath()

	return &GetIndexPageDataOutput{
		PageData: indexPage,
		Template: htmlTemplateFullPath,
		Error:    err,
	}
}
