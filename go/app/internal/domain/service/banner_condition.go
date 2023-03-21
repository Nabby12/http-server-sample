package service

import (
	"time"
)

type (
	BannerCondition interface {
		IsBannerDisplayed(input BannerConditionInput) *BannerConditionOutput
	}
	BannerConditionInput struct {
		Location          *time.Location
		CurrentTimeString string
		ClientIP          string
		BannerStartTime   string
		BannerEndTime     string
		BannerTargetIP    string
	}
	BannerConditionOutput struct {
		Result bool
		Err    error
	}
)
