package repository_test

import (
	"log"
	"testing"
	"time"

	"go-http-server/interface/repository"
)

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

func TestIsShowBanner(t *testing.T) {
	// テスト用環境変数
	t.Setenv("TARGET_IP", "999.999.999.999")

	// test用データ
	testData := []struct {
		clientIP    string
		currentTime time.Time
		startTime   time.Time
		endTime     time.Time
		want        bool
	}{
		{
			// 指定期間内でバナー表示
			clientIP:    "111.111.111.111",
			currentTime: setTestTime("2022/06/12 09:00:00.000"),
			startTime:   setTestTime("2022/06/11 09:00:00.000"),
			endTime:     setTestTime("2022/06/13 09:00:00.000"),
			want:        true,
		},
		{
			// 開始期間前のためバナー非表示
			clientIP:    "111.111.111.111",
			currentTime: setTestTime("2022/06/11 09:00:00.000"),
			startTime:   setTestTime("2022/06/12 09:00:00.000"),
			endTime:     setTestTime("2022/06/13 09:00:00.000"),
			want:        false,
		},
		{
			// 終了期間後のためバナー非表示
			clientIP:    "111.111.111.111",
			currentTime: setTestTime("2022/06/13 09:00:00.000"),
			startTime:   setTestTime("2022/06/11 09:00:00.000"),
			endTime:     setTestTime("2022/06/12 09:00:00.000"),
			want:        false,
		},
		{
			// 指定ipのため表示期間前だがバナー表示
			clientIP:    "999.999.999.999",
			currentTime: setTestTime("2022/06/11 09:00:00.000"),
			startTime:   setTestTime("2022/06/12 09:00:00.000"),
			endTime:     setTestTime("2022/06/13 09:00:00.000"),
			want:        true,
		},
		{
			// 指定ipだが、表示期間後のためバナー非表示
			clientIP:    "999.999.999.999",
			currentTime: setTestTime("2022/06/13 09:00:00.000"),
			startTime:   setTestTime("2022/06/11 09:00:00.000"),
			endTime:     setTestTime("2022/06/12 09:00:00.000"),
			want:        false,
		},
	}

	for _, td := range testData {
		result, err := repository.IsShowBanner(
			td.clientIP,
			td.currentTime,
			td.startTime,
			td.endTime,
		)
		if err != nil {
			t.Error(err)
		}

		if result != td.want {
			t.Error("the result is not expected.")
		}
	}
}
