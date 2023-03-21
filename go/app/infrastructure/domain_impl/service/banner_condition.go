package service

import (
	"go-http-server/internal/domain/service"
	"time"
)

type bannerConditionImpl struct{}

func NewBannerCondition() service.BannerCondition {
	return &bannerConditionImpl{}
}

func (r *bannerConditionImpl) IsBannerDisplayed(input service.BannerConditionInput) *service.BannerConditionOutput {
	currentTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", input.CurrentTimeString, input.Location)
	if err != nil {
		return &service.BannerConditionOutput{
			Result: false,
			Err:    err,
		}
	}
	startTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", input.BannerStartTime, input.Location)
	if err != nil {
		return &service.BannerConditionOutput{
			Result: false,
			Err:    err,
		}
	}
	endTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", input.BannerEndTime, input.Location)
	if err != nil {
		return &service.BannerConditionOutput{
			Result: false,
			Err:    err,
		}
	}

	// startTime <= 現在時刻 <= endTime の場合バナー表示
	if !startTime.After(currentTime) && !currentTime.After(endTime) {
		return &service.BannerConditionOutput{
			Result: true,
			Err:    nil,
		}
	}

	// 「特定IPからのアクセス」かつ「現在時刻 <= endTime」時もバナー表示
	if input.ClientIP == input.BannerTargetIP && !currentTime.After(endTime) {
		return &service.BannerConditionOutput{
			Result: true,
			Err:    nil,
		}
	}

	return &service.BannerConditionOutput{
		Result: false,
		Err:    nil,
	}
}
