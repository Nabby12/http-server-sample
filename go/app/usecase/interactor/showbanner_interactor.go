package interactor

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"go-http-server/usecase/repository"
)

type showbannerInteractor struct {
	Repository repository.ShowbannerRepository
}

type ShowbannerInteractor interface {
	GetView(http.ResponseWriter, *http.Request) error
}

func NewShowbannerInteractor(sr repository.ShowbannerRepository) ShowbannerInteractor {
	return &showbannerInteractor{
		Repository: sr,
	}
}

func (si *showbannerInteractor) GetView(w http.ResponseWriter, r *http.Request) error {
	// timezone を設定
	const localTimeZone string = "Asia/Tokyo"
	localLocation := time.FixedZone(localTimeZone, 9*60*60)

	currentTime := time.Now().In(localLocation)

	startTimeEnv, ok := os.LookupEnv("START_TIME")
	if !ok {
		startTimeEnv = "2006/01/02 15:04:05.000"
	}
	startTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", startTimeEnv, localLocation)
	if err != nil {
		fmt.Printf("Failed Set starttime. due to an error: %v\n", err)
		return err
	}

	endTimeEnv, ok := os.LookupEnv("END_TIME")
	if !ok {
		endTimeEnv = "2006/01/02 15:04:05.000"
	}
	endTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", endTimeEnv, localLocation)
	if err != nil {
		fmt.Printf("Failed Set endtime. due to an error: %v\n", err)
		return err
	}

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

	if err := si.Repository.SetView(w, clientIP, currentTime, startTime, endTime); err != nil {
		return err
	}

	return nil
}
