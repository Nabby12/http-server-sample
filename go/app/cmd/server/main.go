package main

import (
	"log"
	"net/http"

	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/adapter/registry"
	"go-http-server/internal/adapter/router"
)

func main() {
	// 環境変数読み込みもここかな
	configuration.SetEnv()

	mux := http.NewServeMux()
	registry := registry.Initialize()
	router.Route(mux, registry)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed Listen and Serve. due to an error: ", err)
	}
}
