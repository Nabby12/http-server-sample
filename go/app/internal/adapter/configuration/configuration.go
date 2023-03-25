package configuration

type Config struct {
	PublicPath      string
	ImagePath       string
	Location        string
	IndexPage       IndexPage
	ShowBannerPage  ShowBannerPage
	BannerCondition BannerCondition
}

type IndexPage struct {
	Title  string
	Header string
}

type ShowBannerPage struct {
	Title  string
	Header string
}

type BannerCondition struct {
	TagetIP   string
	StartTime string
	EndTime   string
}

var config *Config

func SetEnv() {
	config = &Config{
		PublicPath: "/go/src/app/infrastructure/public",
		ImagePath:  "infrastructure/public",
		Location:   "Asia/Tokyo",
		IndexPage: IndexPage{
			Title:  "Banner Sample via Golang",
			Header: "Index",
		},
		ShowBannerPage: ShowBannerPage{
			Title:  "Banner Sample via Golang",
			Header: "Show Banner",
		},
		BannerCondition: BannerCondition{
			TagetIP:   "xx.xx.xx.xx",
			StartTime: "1900/01/01 00:00:00.000",
			EndTime:   "2099/01/01 00:00:00.000",
		},
	}
}

func LoadEnv() Config {
	return *config
}
