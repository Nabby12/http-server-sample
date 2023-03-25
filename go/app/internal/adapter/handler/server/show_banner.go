package server

import (
	"go-http-server/internal/adapter/configuration"
	"go-http-server/internal/usecase"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type ShowBannerHandler struct {
	getShowBannerPageData usecase.GetShowBannerPageData
}

func NewShowBannerHandler(getShowBannerPageData usecase.GetShowBannerPageData) *ShowBannerHandler {
	return &ShowBannerHandler{
		getShowBannerPageData: getShowBannerPageData,
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
	targetIP := envValues.BannerCondition.TagetIP

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

	input := usecase.GetShowBannerPageDataInput{
		Title:       title,
		Header:      header,
		CurrentTime: currentTime,
		StartTime:   startTime,
		EndTime:     endTime,
		Location:    location,
		ClientIP:    clientIP,
		TargetIP:    targetIP,
	}

	output := h.getShowBannerPageData.Execute(input)
	if output.Error != nil {
		log.Fatal(output.Error)
	}

	t, err := template.ParseFiles(output.Template)
	if err != nil {
		log.Fatalf("Failed Create template. due to an error: %v\n", err)
	}

	if err := t.Execute(responseWriter, output.PageData); err != nil {
		log.Fatalf("Failed Execute template. due to an error: %v\n", err)
	}
}
