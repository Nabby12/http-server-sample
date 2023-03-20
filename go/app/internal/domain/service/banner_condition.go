package service

type BannerCondition interface {
	IsBannerDisplayed(currentTimeString string, clientIP string) (bool, error)
}
