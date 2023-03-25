package model

import "go-http-server/internal/domain/model"

type showBannerPage struct {
	title       string
	header      string
	currentTime string
	startTime   string
	endTime     string
	bannerFlag  bool
}

func NewShowBannerPage(
	title string,
	header string,
	currentTime string,
	startTime string,
	endTime string,
	bannerFlag bool,
) model.ShowBannerPage {
	return &showBannerPage{
		title:       title,
		header:      header,
		currentTime: currentTime,
		startTime:   startTime,
		endTime:     endTime,
		bannerFlag:  bannerFlag,
	}
}

func (sbp *showBannerPage) Title() string {
	return sbp.title
}

func (sbp *showBannerPage) Header() string {
	return sbp.header
}

func (sbp *showBannerPage) CurrentTime() string {
	return sbp.currentTime
}

func (sbp *showBannerPage) StartTime() string {
	return sbp.startTime
}

func (sbp *showBannerPage) EndTime() string {
	return sbp.endTime
}

func (sbp *showBannerPage) BannerFlag() bool {
	return sbp.bannerFlag
}
