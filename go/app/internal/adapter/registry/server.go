package registry

import (
	"go-http-server/infrastructure/domain_impl/repository"
	"go-http-server/infrastructure/domain_impl/service"
	"go-http-server/internal/adapter/handler/server"
	"go-http-server/internal/usecase"
)

type ServerRegistry struct {
	IndexHandler      *server.IndexHandler
	ShowBannerHandler *server.ShowBannerHandler
}

func Initialize() *ServerRegistry {
	htmlRepository := repository.NewHtmlTemplate()
	getIndexPageData := usecase.NewGetIndexPageData(htmlRepository)
	bannerConditionService := service.NewBannerCondition()
	getShowBannerPageData := usecase.NewGetShowBannerPageData(htmlRepository, bannerConditionService)
	return &ServerRegistry{
		IndexHandler:      server.NewIndexHandler(getIndexPageData),
		ShowBannerHandler: server.NewShowBannerHandler(getShowBannerPageData),
	}
}
