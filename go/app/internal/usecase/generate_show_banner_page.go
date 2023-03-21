package usecase

import (
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"go-http-server/internal/domain/service"
	"net/http"
	"time"
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
		ShowBannerPage  model.ShowBannerPage
		ResponseWriter  http.ResponseWriter
		Location        *time.Location
		ClientIP        string
		BannerStartTime string
		BannerEndTime   string
		BannerTargetIP  string
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
	bannerConditionInput := service.BannerConditionInput{
		Location:          input.Location,
		CurrentTimeString: input.ShowBannerPage.CurrentTime(),
		ClientIP:          input.ClientIP,
		BannerStartTime:   input.BannerStartTime,
		BannerEndTime:     input.BannerEndTime,
		BannerTargetIP:    input.BannerTargetIP,
	}
	isBannerDisplayedOutput := u.bannerConditionService.IsBannerDisplayed(bannerConditionInput)
	if isBannerDisplayedOutput.Err != nil {
		return &GenerateShowBannerPageOutput{
			Error: isBannerDisplayedOutput.Err,
		}
	}

	// バナー表示フラグを更新
	input.ShowBannerPage.UpdateBannerFlag(isBannerDisplayedOutput.Result)

	if err := u.showBannerPageRepository.SetView(input.ShowBannerPage, input.ResponseWriter); err != nil {
		return &GenerateShowBannerPageOutput{
			Error: err,
		}
	}

	return &GenerateShowBannerPageOutput{
		Error: nil,
	}
}
