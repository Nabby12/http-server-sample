package usecase

import (
	inframodel "go-http-server/infrastructure/domain_impl/model"
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"go-http-server/internal/domain/service"
	"time"
)

type (
	getShowBannerPageDataImpl struct {
		HtmlTemplateRepository repository.HtmlTemplate
		bannerConditionService service.BannerCondition
	}
	GetShowBannerPageData interface {
		Execute(input GetShowBannerPageDataInput) *GetShowBannerPageDataOutput
	}
	GetShowBannerPageDataInput struct {
		Title       string
		Header      string
		CurrentTime string
		StartTime   string
		EndTime     string
		Location    *time.Location
		ClientIP    string
		TargetIP    string
	}

	GetShowBannerPageDataOutput struct {
		PageData model.ShowBannerPage
		Template string
		Error    error
	}
)

var (
	ShowBannerKey = "showbanner"
)

func NewGetShowBannerPageData(
	HtmlTemplateRepository repository.HtmlTemplate,
	bannerConditionService service.BannerCondition,
) GetShowBannerPageData {
	return &getShowBannerPageDataImpl{
		HtmlTemplateRepository: HtmlTemplateRepository,
		bannerConditionService: bannerConditionService,
	}
}

func (u *getShowBannerPageDataImpl) Execute(input GetShowBannerPageDataInput) *GetShowBannerPageDataOutput {
	// バナー表示判定
	bannerConditionInput := service.BannerConditionInput{
		Location:          input.Location,
		CurrentTimeString: input.CurrentTime,
		ClientIP:          input.ClientIP,
		BannerStartTime:   input.StartTime,
		BannerEndTime:     input.EndTime,
		BannerTargetIP:    input.TargetIP,
	}
	isBannerDisplayedOutput := u.bannerConditionService.IsBannerDisplayed(bannerConditionInput)
	if isBannerDisplayedOutput.Err != nil {
		return &GetShowBannerPageDataOutput{
			Error: isBannerDisplayedOutput.Err,
		}
	}

	showBannerPage := inframodel.NewShowBannerPage(
		input.Title,
		input.Header,
		input.CurrentTime,
		input.StartTime,
		input.EndTime,
		isBannerDisplayedOutput.Result,
	)

	htmlTemplate, err := u.HtmlTemplateRepository.GetByName(ShowBannerKey)
	if err != nil {
		return &GetShowBannerPageDataOutput{
			PageData: nil,
			Template: "",
			Error:    err,
		}
	}
	htmlTemplateFullPath := htmlTemplate.FullPath()

	return &GetShowBannerPageDataOutput{
		PageData: showBannerPage,
		Template: htmlTemplateFullPath,
		Error:    err,
	}
}
