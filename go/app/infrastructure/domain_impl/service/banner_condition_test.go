package service

import (
	"go-http-server/internal/domain/service"
	"log"
	"testing"
	"time"
)

func Test_IsBannerDisplayed(t *testing.T) {
	var (
		dummyNoMatchClientIP    = "111.111.111.111"
		dummyMatchClientIP      = "999.999.999.999"
		dummyTargetIP           = "999.999.999.999"
		dummyTimeStringYear2000 = "2000/01/01 09:00:00.000"
		dummyTimeStringYear1950 = "1950/01/01 09:00:00.000"
		dummyTimeStringYear2050 = "2050/01/01 09:00:00.000"
	)

	dummyLocation, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("Failed Setting location. due to an error: ", err)
	}

	tests := []struct {
		name    string
		input   service.BannerConditionInput
		wantRes bool
		wantErr error
	}{
		{
			name: "show banner by between start and end",
			input: service.BannerConditionInput{
				Location:          dummyLocation,
				CurrentTimeString: dummyTimeStringYear2000,
				ClientIP:          dummyNoMatchClientIP,
				BannerStartTime:   dummyTimeStringYear1950,
				BannerEndTime:     dummyTimeStringYear2050,
				BannerTargetIP:    dummyTargetIP,
			},
			wantRes: true,
			wantErr: nil,
		},
		{
			name: "not show banner current time is after end time",
			input: service.BannerConditionInput{
				Location:          dummyLocation,
				CurrentTimeString: dummyTimeStringYear2050,
				ClientIP:          dummyNoMatchClientIP,
				BannerStartTime:   dummyTimeStringYear1950,
				BannerEndTime:     dummyTimeStringYear2000,
				BannerTargetIP:    dummyTargetIP,
			},
			wantRes: false,
			wantErr: nil,
		},
		{
			name: "not show banner current time is before start time",
			input: service.BannerConditionInput{
				Location:          dummyLocation,
				CurrentTimeString: dummyTimeStringYear1950,
				ClientIP:          dummyNoMatchClientIP,
				BannerStartTime:   dummyTimeStringYear2000,
				BannerEndTime:     dummyTimeStringYear2050,
				BannerTargetIP:    dummyTargetIP,
			},
			wantRes: false,
			wantErr: nil,
		},
		{
			name: "show banner by between start and end with matching IP",
			input: service.BannerConditionInput{
				Location:          dummyLocation,
				CurrentTimeString: dummyTimeStringYear2000,
				ClientIP:          dummyMatchClientIP,
				BannerStartTime:   dummyTimeStringYear1950,
				BannerEndTime:     dummyTimeStringYear2050,
				BannerTargetIP:    dummyTargetIP,
			},
			wantRes: true,
			wantErr: nil,
		},
		{
			name: "show banner current time is before end time and matching IP",
			input: service.BannerConditionInput{
				Location:          dummyLocation,
				CurrentTimeString: dummyTimeStringYear1950,
				ClientIP:          dummyMatchClientIP,
				BannerStartTime:   dummyTimeStringYear2000,
				BannerEndTime:     dummyTimeStringYear2050,
				BannerTargetIP:    dummyTargetIP,
			},
			wantRes: true,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewBannerCondition()

			output := service.IsBannerDisplayed(tt.input)
			if output.Result != tt.wantRes || output.Err != tt.wantErr {
				t.Fatal("failed: result or error is invalid")
			}
		})
	}
}
