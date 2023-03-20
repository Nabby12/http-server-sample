package service

import (
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/domain/service"
	"time"
)

type bannerConditionImpl struct{}

func NewBannerCondition() service.BannerCondition {
	return &bannerConditionImpl{}
}

func (r *bannerConditionImpl) IsBannerDisplayed(currentTimeString string, clientIP string) (bool, error) {
	envValues := configuration.LoadEnv()
	location, err := time.LoadLocation(envValues.Location)
	if err != nil {
		return false, err
	}

	currentTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", currentTimeString, location)
	if err != nil {
		return false, err
	}
	startTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", envValues.BannerCondition.StartTime, location)
	if err != nil {
		return false, err
	}
	endTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", envValues.BannerCondition.EndTime, location)
	if err != nil {
		return false, err
	}

	bannerFlag := false
	// startTime <= 現在時刻 <= endTime の場合バナー表示
	if !startTime.After(currentTime) && !currentTime.After(endTime) {
		bannerFlag = true
	}

	targetIP := envValues.BannerCondition.TagetIP

	// 「特定IPからのアクセス」かつ「現在時刻 <= endTime」時もバナー表示
	if clientIP == targetIP && !currentTime.After(endTime) {
		bannerFlag = true
	}

	return bannerFlag, nil
}
