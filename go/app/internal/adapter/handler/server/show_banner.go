package server

import (
	"go-http-server/infrastructure/domain_impl/model"
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/usecase"
	"log"
	"net/http"
	"strings"
	"time"
)

type ShowBannerHandler struct {
	generateShowBannerPage usecase.GenerateShowBannerPage
}

func NewShowBannerHandler(generateShowBannerPage usecase.GenerateShowBannerPage) *ShowBannerHandler {
	return &ShowBannerHandler{
		generateShowBannerPage: generateShowBannerPage,
	}
}

func (h *ShowBannerHandler) Handle(responseWriter http.ResponseWriter, r *http.Request) {
	envValues := configuration.LoadEnv()
	location, err := time.LoadLocation(envValues.Location)
	if err != nil {
		log.Fatal("Failed Setting location. due to an error: ", err)
	}
	currentTimeString := time.Now().In(location).Format("2006/01/02 15:04:05.000")

	title := envValues.ShowBannerPage.Title
	header := envValues.ShowBannerPage.Header
	currentTime := currentTimeString
	startTime := envValues.BannerCondition.StartTime
	endTime := envValues.BannerCondition.EndTime
	bannerFlag := false

	showBannerPage := model.NewShowBannerPage(title, header, currentTime, startTime, endTime, bannerFlag)

	var clientIP string
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		clientIP = forwarded
	} else {
		clientIP = r.RemoteAddr
	}
	if strings.Contains(clientIP, ":") {
		clientIP = strings.Split(clientIP, ":")[0]
	}

	if output := h.generateShowBannerPage.Execute(usecase.GenerateShowBannerPageInput{
		ShowBannerPage: showBannerPage,
		ClientIP:       clientIP,
		ResponseWriter: responseWriter,
	}); output.Error != nil {
		log.Fatal(output.Error)
	}
}
