package usecase

import (
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"go-http-server/internal/domain/service"
	"net/http"
)

type (
	generateShowBannerPageImpl struct {
		showBannerPageRepository repository.ShowBannerPage
		bannerConditionService   service.BannerCondition
	}
	GenerateShowBannerPage interface {
		Execute(input GenerateShowBannerPageInput) *GenerateShowBannerPageOutput
	}
	GenerateShowBannerPageInput struct {
		ShowBannerPage model.ShowBannerPage
		ClientIP       string
		ResponseWriter http.ResponseWriter
	}
	GenerateShowBannerPageOutput struct {
		Error error
	}
)

func NewGenerateShowBannerPage(
	showBannerPageRepository repository.ShowBannerPage,
	bannerConditionService service.BannerCondition,
) GenerateShowBannerPage {
	return &generateShowBannerPageImpl{
		showBannerPageRepository: showBannerPageRepository,
		bannerConditionService:   bannerConditionService,
	}
}

func (u *generateShowBannerPageImpl) Execute(input GenerateShowBannerPageInput) *GenerateShowBannerPageOutput {
	// バナー表示判定
	isBannerDisplayed, err := u.bannerConditionService.IsBannerDisplayed(input.ShowBannerPage.CurrentTime(), input.ClientIP)
	if err != nil {
		return &GenerateShowBannerPageOutput{
			Error: err,
		}
	}

	// バナー表示フラグを更新
	input.ShowBannerPage.UpdateBannerFlag(isBannerDisplayed)

	if err := u.showBannerPageRepository.SetView(input.ShowBannerPage, input.ResponseWriter); err != nil {
		return &GenerateShowBannerPageOutput{
			Error: err,
		}
	}

	return &GenerateShowBannerPageOutput{
		Error: nil,
	}
}
