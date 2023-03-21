package usecase

import (
	"errors"
	"fmt"
	"go-http-server/infrastructure/domain_impl/model"
	domain_model "go-http-server/internal/domain/model"
	"go-http-server/internal/domain/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockShowBannerPageRepository struct {
	setViewCallTimes        int
	UpdateBannerFlagTimes   int
	lastInputShowBannerPage domain_model.ShowBannerPage
	lastInputResponseWriter http.ResponseWriter
	lastInputBannerFlag     bool
	err                     error
}

func (m *mockShowBannerPageRepository) SetView(showBannerPage domain_model.ShowBannerPage, responseWriter http.ResponseWriter) error {
	m.setViewCallTimes++
	m.lastInputShowBannerPage = showBannerPage
	m.lastInputResponseWriter = responseWriter
	return m.err
}
func (m *mockShowBannerPageRepository) UpdateBannerFlag(flag bool) {
	m.setViewCallTimes++
	m.lastInputBannerFlag = flag
}

type mockBannerConditionService struct {
	callTimes int
	lastInput service.BannerConditionInput
	output    *service.BannerConditionOutput
}

func (m *mockBannerConditionService) IsBannerDisplayed(bannerConditionInput service.BannerConditionInput) *service.BannerConditionOutput {
	m.callTimes++
	m.lastInput = bannerConditionInput
	return m.output
}

func Test_GenerateShowBannerPage_Execute(t *testing.T) {
	dummylocation, _ := time.LoadLocation("Asia/Tokyo")
	var (
		dummyTitle              = "dummy title"
		dummyHeader             = "dummy header"
		dummyResponseWriter     = httptest.NewRecorder()
		dummyClientIP           = "111.111.111.111"
		dummyTargetIP           = "999.999.999.999"
		dummyTimeStringYear2000 = "2000/01/01 09:00:00.000"
		dummyTimeStringYear1950 = "1950/01/01 09:00:00.000"
		dummyTimeStringYear2050 = "2050/01/01 09:00:00.000"
		dummyError              = errors.New("dummy error")

		dummyShowBannerPageBannerFlagTrue = model.NewShowBannerPage(
			dummyTitle,
			dummyHeader,
			dummyTimeStringYear2000,
			dummyTimeStringYear1950,
			dummyTimeStringYear2050,
			true,
		)
		dummyShowBannerPageBannerFlagFalse = model.NewShowBannerPage(
			dummyTitle,
			dummyHeader,
			dummyTimeStringYear2000,
			dummyTimeStringYear1950,
			dummyTimeStringYear2050,
			false,
		)

		dummyBannerConditionInput = service.BannerConditionInput{
			Location:          dummylocation,
			CurrentTimeString: dummyTimeStringYear2000,
			ClientIP:          dummyClientIP,
			BannerStartTime:   dummyTimeStringYear1950,
			BannerEndTime:     dummyTimeStringYear2050,
			BannerTargetIP:    dummyTargetIP,
		}
	)

	type (
		input struct {
			title          string
			header         string
			currentTime    string
			startTime      string
			endTime        string
			responseWriter *httptest.ResponseRecorder
			location       *time.Location
			clientIP       string
			bannerTargetIP string
		}
		mock struct {
			repoErr           error
			isBannerDisplayed bool
			serviceErr        error
		}
		want struct {
			setViewCallTimes           int
			UpdateBannerFlagTimes      int
			callInputShowBannerPage    domain_model.ShowBannerPage
			callInputResponseWriter    http.ResponseWriter
			callInputBannerFlag        bool
			isBannerDisplayedCallTimes int
			callBannerConditionInput   service.BannerConditionInput
			bannerConditionOutput      *service.BannerConditionOutput
		}
	)

	setup := func(input input, mock mock) (GenerateShowBannerPageInput, *mockShowBannerPageRepository, *mockBannerConditionService) {
		showBannerPage := model.NewShowBannerPage(
			input.title,
			input.header,
			input.currentTime,
			input.startTime,
			input.endTime,
			false,
		)
		usecaseInput := GenerateShowBannerPageInput{
			ShowBannerPage:  showBannerPage,
			ResponseWriter:  input.responseWriter,
			Location:        input.location,
			ClientIP:        input.clientIP,
			BannerStartTime: input.startTime,
			BannerEndTime:   input.endTime,
			BannerTargetIP:  input.bannerTargetIP,
		}

		mockShowBannerPageRepository := &mockShowBannerPageRepository{
			err: mock.repoErr,
		}

		mockBannerConditionService := &mockBannerConditionService{
			output: &service.BannerConditionOutput{
				Result: mock.isBannerDisplayed,
				Err:    mock.serviceErr,
			},
		}

		return usecaseInput, mockShowBannerPageRepository, mockBannerConditionService
	}

	tests := []struct {
		name    string
		input   input
		mock    mock
		want    want
		wantErr error
	}{
		{
			name: "success: no error with bannerFlag true",
			input: input{
				title:          dummyTitle,
				header:         dummyHeader,
				currentTime:    dummyTimeStringYear2000,
				startTime:      dummyTimeStringYear1950,
				endTime:        dummyTimeStringYear2050,
				responseWriter: dummyResponseWriter,
				location:       dummylocation,
				clientIP:       dummyClientIP,
				bannerTargetIP: dummyTargetIP,
			},
			mock: mock{
				repoErr:           nil,
				isBannerDisplayed: true,
				serviceErr:        nil,
			},
			want: want{
				setViewCallTimes:           1,
				UpdateBannerFlagTimes:      1,
				callInputShowBannerPage:    dummyShowBannerPageBannerFlagTrue,
				callInputResponseWriter:    dummyResponseWriter,
				callInputBannerFlag:        true,
				isBannerDisplayedCallTimes: 1,
				callBannerConditionInput:   dummyBannerConditionInput,
				bannerConditionOutput: &service.BannerConditionOutput{
					Result: true,
				},
			},
			wantErr: nil,
		},
		{
			name: "success: no error with bannerFlag false",
			input: input{
				title:          dummyTitle,
				header:         dummyHeader,
				currentTime:    dummyTimeStringYear2000,
				startTime:      dummyTimeStringYear1950,
				endTime:        dummyTimeStringYear2050,
				responseWriter: dummyResponseWriter,
				location:       dummylocation,
				clientIP:       dummyClientIP,
				bannerTargetIP: dummyTargetIP,
			},
			mock: mock{
				repoErr:           nil,
				isBannerDisplayed: false,
				serviceErr:        nil,
			},
			want: want{
				setViewCallTimes:           1,
				UpdateBannerFlagTimes:      1,
				callInputShowBannerPage:    dummyShowBannerPageBannerFlagFalse,
				callInputResponseWriter:    dummyResponseWriter,
				callInputBannerFlag:        true,
				isBannerDisplayedCallTimes: 1,
				callBannerConditionInput:   dummyBannerConditionInput,
				bannerConditionOutput: &service.BannerConditionOutput{
					Result: false,
				},
			},
			wantErr: nil,
		},
		{
			name: "failure: repository setview error and return error",
			input: input{
				title:          dummyTitle,
				header:         dummyHeader,
				currentTime:    dummyTimeStringYear2000,
				startTime:      dummyTimeStringYear1950,
				endTime:        dummyTimeStringYear2050,
				responseWriter: dummyResponseWriter,
				location:       dummylocation,
				clientIP:       dummyClientIP,
				bannerTargetIP: dummyTargetIP,
			},
			mock: mock{
				repoErr:           dummyError,
				isBannerDisplayed: true,
				serviceErr:        nil,
			},
			want: want{
				setViewCallTimes:           1,
				UpdateBannerFlagTimes:      1,
				callInputShowBannerPage:    dummyShowBannerPageBannerFlagTrue,
				callInputResponseWriter:    dummyResponseWriter,
				callInputBannerFlag:        true,
				isBannerDisplayedCallTimes: 1,
				callBannerConditionInput:   dummyBannerConditionInput,
				bannerConditionOutput: &service.BannerConditionOutput{
					Result: true,
				},
			},
			wantErr: dummyError,
		},
		{
			name: "failure: service error and return error",
			input: input{
				title:          dummyTitle,
				header:         dummyHeader,
				currentTime:    dummyTimeStringYear2000,
				startTime:      dummyTimeStringYear1950,
				endTime:        dummyTimeStringYear2050,
				responseWriter: dummyResponseWriter,
				location:       dummylocation,
				clientIP:       dummyClientIP,
				bannerTargetIP: dummyTargetIP,
			},
			mock: mock{
				repoErr:           nil,
				isBannerDisplayed: false,
				serviceErr:        dummyError,
			},
			want: want{
				setViewCallTimes:           1,
				UpdateBannerFlagTimes:      0,
				callInputShowBannerPage:    dummyShowBannerPageBannerFlagFalse,
				callInputResponseWriter:    dummyResponseWriter,
				callInputBannerFlag:        false,
				isBannerDisplayedCallTimes: 1,
				callBannerConditionInput:   dummyBannerConditionInput,
				bannerConditionOutput: &service.BannerConditionOutput{
					Result: false,
				},
			},
			wantErr: dummyError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecaseInput, mockShowBannerPageRepository, mockBannerConditionService := setup(tt.input, tt.mock)
			usecase := NewGenerateShowBannerPage(
				mockShowBannerPageRepository,
				mockBannerConditionService,
			)

			output := usecase.Execute(usecaseInput)
			if mockShowBannerPageRepository.setViewCallTimes != tt.want.setViewCallTimes &&
				mockShowBannerPageRepository.UpdateBannerFlagTimes != tt.want.UpdateBannerFlagTimes {
				t.Fatal("failed: repository call times is invalid")
			}
			// serviceがエラーの場合、後続のSetViewメソッドが呼び出されず、nil pointer error が発生
			if tt.mock.serviceErr == nil {
				if err := AssertEqualShowBannerPageRepository(
					mockShowBannerPageRepository.lastInputShowBannerPage,
					mockShowBannerPageRepository.lastInputResponseWriter,
					tt.want.callInputShowBannerPage,
					tt.want.callInputResponseWriter,
				); err != nil {
					t.Fatal(err.Error())
				}
			}

			if mockBannerConditionService.callTimes != tt.want.isBannerDisplayedCallTimes {
				t.Fatal("failed: service call times is invalid")
			}
			if err := AssertEqualBannerConditionService(
				mockBannerConditionService.lastInput,
				mockBannerConditionService.output,
				tt.want.callBannerConditionInput,
				tt.want.bannerConditionOutput,
			); err != nil {
				t.Fatal(err.Error())
			}

			if output.Error != tt.wantErr {
				t.Fatal("failed: want error is invalid")
			}
		})
	}
}

func AssertEqualShowBannerPageRepository(
	gotShowBannerPage domain_model.ShowBannerPage,
	gotResponseWriter http.ResponseWriter,
	wantShowBannerPage domain_model.ShowBannerPage,
	wantResponseWriter http.ResponseWriter,
) error {
	if gotShowBannerPage.Title() != wantShowBannerPage.Title() ||
		gotShowBannerPage.Header() != wantShowBannerPage.Header() ||
		gotShowBannerPage.CurrentTime() != wantShowBannerPage.CurrentTime() ||
		gotShowBannerPage.StartTime() != wantShowBannerPage.StartTime() ||
		gotShowBannerPage.EndTime() != wantShowBannerPage.EndTime() ||
		gotShowBannerPage.BannerFlag() != wantShowBannerPage.BannerFlag() ||
		gotResponseWriter != wantResponseWriter {
		gotValues := fmt.Sprintf(
			"title: %v, header: %v, currentTime: %v, startTime: %v, endTime: %v, bannerFlag: %v, responseWriter: %v\n",
			gotShowBannerPage.Title(),
			gotShowBannerPage.Header(),
			gotShowBannerPage.CurrentTime(),
			gotShowBannerPage.StartTime(),
			gotShowBannerPage.EndTime(),
			gotShowBannerPage.BannerFlag(),
			gotResponseWriter,
		)
		wantValues := fmt.Sprintf(
			"title: %v, header: %v, currentTime: %v, startTime: %v, endTime: %v, bannerFlag: %v, responseWriter: %v\n",
			wantShowBannerPage.Title(),
			wantShowBannerPage.Header(),
			wantShowBannerPage.CurrentTime(),
			wantShowBannerPage.StartTime(),
			wantShowBannerPage.EndTime(),
			wantShowBannerPage.BannerFlag(),
			wantResponseWriter,
		)

		errMessage := fmt.Sprintf("failed: repository call input is invalid\n  Got : %v \n  Want: %v", gotValues, wantValues)
		return errors.New(errMessage)
	}

	return nil
}

func AssertEqualBannerConditionService(
	gotInput service.BannerConditionInput,
	gotOutput *service.BannerConditionOutput,
	wantInput service.BannerConditionInput,
	wantOutput *service.BannerConditionOutput,
) error {
	if gotInput.Location != wantInput.Location ||
		gotInput.CurrentTimeString != wantInput.CurrentTimeString ||
		gotInput.ClientIP != wantInput.ClientIP ||
		gotInput.BannerStartTime != wantInput.BannerStartTime ||
		gotInput.BannerEndTime != wantInput.BannerEndTime ||
		gotInput.BannerTargetIP != wantInput.BannerTargetIP ||
		gotOutput.Result != wantOutput.Result {
		gotValues := fmt.Sprintf(
			"location: %v, currentTime: %v, clientIP: %v, startTime: %v, endTime: %v, bannerTargetIP: %v, result: %v\n",
			gotInput.Location,
			gotInput.CurrentTimeString,
			gotInput.ClientIP,
			gotInput.BannerStartTime,
			gotInput.BannerEndTime,
			gotInput.BannerTargetIP,
			gotOutput.Result,
		)
		wantValues := fmt.Sprintf(
			"location: %v, currentTime: %v, clientIP: %v, startTime: %v, endTime: %v, bannerTargetIP: %v, result: %v\n",
			wantInput.Location,
			wantInput.CurrentTimeString,
			wantInput.ClientIP,
			wantInput.BannerStartTime,
			wantInput.BannerEndTime,
			wantInput.BannerTargetIP,
			wantOutput.Result,
		)
		errMessage := fmt.Sprintf("failed: service call input is invalid\n  Got : %v \n  Want: %v", gotValues, wantValues)
		return errors.New(errMessage)
	}
	return nil
}
