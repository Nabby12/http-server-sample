package router

import (
	"fmt"
	"net/http"

	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/adapter/registry"
)

func Route(mux *http.ServeMux, registry *registry.ServerRegistry) {
	envValues := configuration.LoadEnv()
	imagePath := fmt.Sprintf("/%v/", envValues.ImagePath)
	imagePathPrefix := fmt.Sprintf("/%v", envValues.ImagePath)
	fileServerPath := envValues.ImagePath
	mux.Handle(imagePath, http.StripPrefix(imagePathPrefix, http.FileServer(http.Dir(fileServerPath))))

	mux.HandleFunc("/", registry.IndexHandler.Handle)
	mux.HandleFunc("/showbanner", registry.ShowBannerHandler.Handle)
}
