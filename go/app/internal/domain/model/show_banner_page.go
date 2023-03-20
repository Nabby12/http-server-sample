package model

type ShowBannerPage interface {
	Title() string
	Header() string
	CurrentTime() string
	StartTime() string
	EndTime() string
	BannerFlag() bool
	UpdateBannerFlag(flag bool)
}
