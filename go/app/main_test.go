package main_test

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"go-http-server/interface/controller"
	interfaceRepository "go-http-server/interface/repository"
	"go-http-server/usecase/interactor"
	usecaseRepository "go-http-server/usecase/repository"
)

func newDummyHandler() *template.Template {
	var template *template.Template
	return template
}

func TestIndexView(t *testing.T) {
	dummyController := controller.NewIndexController(newDummyHandler())

	testserver := httptest.NewServer(http.HandlerFunc(dummyController.View))
	defer testserver.Close()

	res, err := http.Get(testserver.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error("a response code is not 200")
	}

	if !strings.Contains(string(body), "<h1>Index</h1>") {
		t.Error("a response does not contain 'Index' header")
	}
}

// show banner test用の mock 関数等を用意
type showbannerController struct {
	Interactor interactor.ShowbannerInteractor
}

func (sc *showbannerController) View(w http.ResponseWriter, r *http.Request) {
	if err := sc.Interactor.GetView(w, r); err != nil {
		log.Fatal(err)
	}
}

type showbannerInteractorMock1 struct {
	Repository usecaseRepository.ShowbannerRepository
}
type showbannerInteractorMock2 struct {
	Repository usecaseRepository.ShowbannerRepository
}

func newShowbannerInteractorMock1() interactor.ShowbannerInteractor {
	sr := interfaceRepository.NewShowbannerRepository(newDummyHandler())
	return &showbannerInteractorMock1{
		Repository: sr,
	}
}
func newShowbannerInteractorMock2() interactor.ShowbannerInteractor {
	sr := interfaceRepository.NewShowbannerRepository(newDummyHandler())
	return &showbannerInteractorMock2{
		Repository: sr,
	}
}

func (si *showbannerInteractorMock1) GetView(w http.ResponseWriter, r *http.Request) error {
	// 指定期間内でバナー表示
	clientIP := "111.111.111.111"
	currentTime := setTestTime("2022/06/12 09:00:00.000")
	startTime := setTestTime("2022/06/11 09:00:00.000")
	endTime := setTestTime("2022/06/13 09:00:00.000")

	if err := si.Repository.SetView(w, clientIP, currentTime, startTime, endTime); err != nil {
		return err
	}
	return nil
}
func (si *showbannerInteractorMock2) GetView(w http.ResponseWriter, r *http.Request) error {
	// 開始期間外でバナー非表示
	clientIP := "111.111.111.111"
	currentTime := setTestTime("2022/06/11 09:00:00.000")
	startTime := setTestTime("2022/06/12 09:00:00.000")
	endTime := setTestTime("2022/06/13 09:00:00.000")

	if err := si.Repository.SetView(w, clientIP, currentTime, startTime, endTime); err != nil {
		return err
	}
	return nil
}

func setTestTime(timeString string) time.Time {
	// timezone を設定
	const localTimeZone string = "Asia/Tokyo"
	localLocation := time.FixedZone(localTimeZone, 9*60*60)

	testTime, err := time.ParseInLocation("2006/01/02 15:04:05.000", timeString, localLocation)
	if err != nil {
		log.Fatal("Failed Setting starttime. due to an error: ", err)
	}

	return testTime
}

func TestShowbannerView(t *testing.T) {
	// テスト用環境変数
	t.Setenv("TARGET_IP", "999.999.999.999")

	// test用データ
	testData := []struct {
		newSiMock  interactor.ShowbannerInteractor
		bannerFlag bool
	}{
		{
			// 指定期間内でバナー表示
			newShowbannerInteractorMock1(),
			true,
		},
		{
			// 開始期間外でバナー非表示
			newShowbannerInteractorMock2(),
			false,
		},
	}

	for _, td := range testData {
		siMock := td.newSiMock
		dummyController := &showbannerController{
			Interactor: siMock,
		}

		testserver := httptest.NewServer(http.HandlerFunc(dummyController.View))
		defer testserver.Close()

		res, err := http.Get(testserver.URL)
		if err != nil {
			t.Error(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 200 {
			t.Error("a response code is not 200")
		}

		if !strings.Contains(string(body), "<h1>Show Banner</h1>") {
			t.Error("a response does not contain 'Show Banner' header")
		}

		// バナー表示が正しくできているか確認
		const bannerElement string = "<img src=\"/static/banner.png\" alt=\"sample banner\" height=\"200\">"
		if td.bannerFlag {
			if !strings.Contains(string(body), bannerElement) {
				t.Error("a response does not contain banner element")
			}
		} else {
			if strings.Contains(string(body), bannerElement) {
				t.Error("a response contains banner element")
			}
		}

	}
}
