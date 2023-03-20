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
	indexRepository := repository.NewIndexPage()
	generateIndexPage := usecase.NewGenerateIndexPage(indexRepository)
	showBannerRepository := repository.NewShowBannerPage()
	bannerConditionService := service.NewBannerCondition()
	generateShowBannerPage := usecase.NewGenerateShowBannerPage(showBannerRepository, bannerConditionService)
	return &ServerRegistry{
		IndexHandler:      server.NewIndexHandler(generateIndexPage),
		ShowBannerHandler: server.NewShowBannerHandler(generateShowBannerPage),
	}
}
