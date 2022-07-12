package repository

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"go-http-server/domain"
	"go-http-server/usecase/repository"
)

type showbannerRepository struct {
	Template *template.Template
}

func NewShowbannerRepository(t *template.Template) repository.ShowbannerRepository {
	return &showbannerRepository{
		Template: t,
	}
}

func (sr *showbannerRepository) SetView(
	w http.ResponseWriter,
	clientIP string,
	currentTime time.Time,
	startTime time.Time,
	endTime time.Time,
) error {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed Get current path. due to an error: %v\n", err)
		return err
	}

	bannerFlag, err := IsShowBanner(
		clientIP,
		currentTime,
		startTime,
		endTime,
	)
	if err != nil {
		return err
	}

	// 画面に表示する日付形式を指定
	timeFormat := "2006-01-02 15:04:05"

	data := domain.ShowBannerPage{
		Title:       "Banner Sample via Golang",
		Header:      "Show Banner",
		CurrentTime: currentTime.Format(timeFormat),
		StartTime:   startTime.Format(timeFormat),
		EndTime:     endTime.Format(timeFormat),
		BannerFlag:  bannerFlag,
	}

	t, err := sr.Template.ParseFiles(currentPath + "/static/showbanner.html")
	if err != nil {
		fmt.Printf("Failed Create template. due to an error: %v\n", err)
		return err
	}

	if err := t.Execute(w, data); err != nil {
		fmt.Printf("Failed Execute template. due to an error: %v\n", err)
		return err
	}

	return nil
}

func IsShowBanner(
	clientIP string,
	currentTime time.Time,
	startTime time.Time,
	endTime time.Time,
) (bool, error) {
	bannerFlag := false

	// startTime <= 現在時刻 <= endTime の場合バナー表示
	if !startTime.After(currentTime) && !currentTime.After(endTime) {
		bannerFlag = true
	}

	targetIP := os.Getenv("TARGET_IP")

	// 「特定IPからのアクセス」かつ「現在時刻 <= endTime」時もバナー表示
	if clientIP == targetIP && !currentTime.After(endTime) {
		bannerFlag = true
	}

	return bannerFlag, nil
}
