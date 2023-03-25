package usecase

import (
	"errors"
	"fmt"
	inframodel "go-http-server/infrastructure/domain_impl/model"
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"go-http-server/internal/domain/service"
	"testing"
	"time"
)

func Test_GetShowBannerPage_Execute(t *testing.T) {
	dummyLocation, _ := time.LoadLocation("Asia/Tokyo")
	var (
		dummyTitle              = "dummy title"
		dummyHeader             = "dummy header"
		dummyKey                = ShowBannerKey
		dummyTemplateFileName   = "dummy html"
		dummyPublicPath         = "dummy public"
		dummyTemplateFullPath   = "dummy public/dummy html"
		dummyClientIP           = "111.111.111.111"
		dummyTargetIP           = "999.999.999.999"
		dummyTimeStringYear2000 = "2000/01/01 09:00:00.000"
		dummyTimeStringYear1950 = "1950/01/01 09:00:00.000"
		dummyTimeStringYear2050 = "2050/01/01 09:00:00.000"
		dummyError              = errors.New("dummy error")

		dummyHtmlPage = inframodel.NewHtmlTemplate(
			dummyTemplateFileName,
			dummyPublicPath,
		)
		dummyShowBannerPageBannerFlagTrue = inframodel.NewShowBannerPage(
			dummyTitle,
			dummyHeader,
			dummyTimeStringYear2000,
			dummyTimeStringYear1950,
			dummyTimeStringYear2050,
			true,
		)
		dummyShowBannerPageBannerFlagFalse = inframodel.NewShowBannerPage(
			dummyTitle,
			dummyHeader,
			dummyTimeStringYear2000,
			dummyTimeStringYear1950,
			dummyTimeStringYear2050,
			false,
		)

		dummyBannerConditionInput = service.BannerConditionInput{
			Location:          dummyLocation,
			CurrentTimeString: dummyTimeStringYear2000,
			ClientIP:          dummyClientIP,
			BannerStartTime:   dummyTimeStringYear1950,
			BannerEndTime:     dummyTimeStringYear2050,
			BannerTargetIP:    dummyTargetIP,
		}
		dummyBannerConditionOutputWithTrue = service.BannerConditionOutput{
			Result: true,
			Err:    nil,
		}
		dummyBannerConditionOutputWithFalse = service.BannerConditionOutput{
			Result: false,
			Err:    nil,
		}
		dummyBannerConditionOutputWithError = service.BannerConditionOutput{
			Result: false,
			Err:    dummyError,
		}
	)

	type (
		input struct {
			title        string
			header       string
			currentTime  string
			startTime    string
			endTime      string
			location     *time.Location
			clientIP     string
			targetIP     string
			serviceInput service.BannerConditionInput
			repoInput    string
		}
		mock struct {
			serviceOutput service.BannerConditionOutput
			repoOutput    model.HtmlTemplate
			repoErr       error
		}
		want struct {
			serviceCallTime int
			serviceInput    service.BannerConditionInput
			serviceOutput   service.BannerConditionOutput
			callTime        int
			callInputKey    string
			callOutput      model.HtmlTemplate
			output          GetShowBannerPageDataOutput
		}
	)

	setup := func(input input, mock mock) (
		*getShowBannerPageDataImpl,
		*repository.MockHtmlTemplateRepository,
		*service.MockBannerConditionService,
	) {
		mockHtmlTemplateRepository := repository.NewMockHtmlTemplateRepository()
		mockHtmlTemplateRepository.EXPECT(
			mock.repoOutput,
			mock.repoErr,
		)
		mockBannerConditionService := service.NewMockBannerConditionService()
		mockBannerConditionService.EXPECT(
			&mock.serviceOutput,
		)

		usecase := &getShowBannerPageDataImpl{
			mockHtmlTemplateRepository,
			mockBannerConditionService,
		}
		return usecase, mockHtmlTemplateRepository, mockBannerConditionService
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
				title:        dummyTitle,
				header:       dummyHeader,
				currentTime:  dummyTimeStringYear2000,
				startTime:    dummyTimeStringYear1950,
				endTime:      dummyTimeStringYear2050,
				location:     dummyLocation,
				clientIP:     dummyClientIP,
				targetIP:     dummyTargetIP,
				serviceInput: dummyBannerConditionInput,
				repoInput:    dummyKey,
			},
			mock: mock{
				serviceOutput: dummyBannerConditionOutputWithTrue,
				repoOutput:    dummyHtmlPage,
				repoErr:       nil,
			},
			want: want{
				serviceCallTime: 1,
				serviceInput:    dummyBannerConditionInput,
				serviceOutput:   dummyBannerConditionOutputWithTrue,
				callTime:        1,
				callInputKey:    dummyKey,
				callOutput:      dummyHtmlPage,
				output: GetShowBannerPageDataOutput{
					PageData: dummyShowBannerPageBannerFlagTrue,
					Template: dummyTemplateFullPath,
					Error:    nil,
				},
			},
		},
		{
			name: "success: no error with bannerFlag false",
			input: input{
				title:        dummyTitle,
				header:       dummyHeader,
				currentTime:  dummyTimeStringYear2000,
				startTime:    dummyTimeStringYear1950,
				endTime:      dummyTimeStringYear2050,
				location:     dummyLocation,
				clientIP:     dummyClientIP,
				targetIP:     dummyTargetIP,
				serviceInput: dummyBannerConditionInput,
				repoInput:    dummyKey,
			},
			mock: mock{
				serviceOutput: dummyBannerConditionOutputWithFalse,
				repoOutput:    dummyHtmlPage,
				repoErr:       nil,
			},
			want: want{
				serviceCallTime: 1,
				serviceInput:    dummyBannerConditionInput,
				serviceOutput:   dummyBannerConditionOutputWithFalse,
				callTime:        1,
				callInputKey:    dummyKey,
				callOutput:      dummyHtmlPage,
				output: GetShowBannerPageDataOutput{
					PageData: dummyShowBannerPageBannerFlagFalse,
					Template: dummyTemplateFullPath,
					Error:    nil,
				},
			},
		},
		{
			name: "failure: repository error and return error",
			input: input{
				title:        dummyTitle,
				header:       dummyHeader,
				currentTime:  dummyTimeStringYear2000,
				startTime:    dummyTimeStringYear1950,
				endTime:      dummyTimeStringYear2050,
				location:     dummyLocation,
				clientIP:     dummyClientIP,
				targetIP:     dummyTargetIP,
				serviceInput: dummyBannerConditionInput,
				repoInput:    dummyKey,
			},
			mock: mock{
				serviceOutput: dummyBannerConditionOutputWithTrue,
				repoOutput:    nil,
				repoErr:       dummyError,
			},
			want: want{
				serviceCallTime: 1,
				serviceInput:    dummyBannerConditionInput,
				serviceOutput:   dummyBannerConditionOutputWithTrue,
				callTime:        1,
				callInputKey:    dummyKey,
				callOutput:      nil,
				output: GetShowBannerPageDataOutput{
					PageData: nil,
					Template: "",
					Error:    dummyError,
				},
			},
		},
		{
			name: "failure: service error and return error",
			input: input{
				title:        dummyTitle,
				header:       dummyHeader,
				currentTime:  dummyTimeStringYear2000,
				startTime:    dummyTimeStringYear1950,
				endTime:      dummyTimeStringYear2050,
				location:     dummyLocation,
				clientIP:     dummyClientIP,
				targetIP:     dummyTargetIP,
				serviceInput: dummyBannerConditionInput,
				repoInput:    dummyKey,
			},
			mock: mock{
				serviceOutput: dummyBannerConditionOutputWithError,
				repoOutput:    nil,
				repoErr:       nil,
			},
			want: want{
				serviceCallTime: 1,
				serviceInput:    dummyBannerConditionInput,
				serviceOutput:   dummyBannerConditionOutputWithError,
				callTime:        0,
				callInputKey:    "",
				callOutput:      nil,
				output: GetShowBannerPageDataOutput{
					PageData: nil,
					Template: "",
					Error:    dummyError,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase, mockRepo, mockService := setup(tt.input, tt.mock)

			output := usecase.Execute(GetShowBannerPageDataInput{
				Title:       tt.input.title,
				Header:      tt.input.header,
				CurrentTime: tt.input.currentTime,
				StartTime:   tt.input.startTime,
				EndTime:     tt.input.endTime,
				Location:    tt.input.location,
				ClientIP:    tt.input.clientIP,
				TargetIP:    tt.input.targetIP,
			})

			if mockRepo.CallTimes != tt.want.callTime {
				t.Fatal("failed: repository call times is invalid")
			}
			if mockRepo.Input != tt.want.callInputKey {
				t.Fatal("failed: repository call input is invalid")
			}
			if mockRepo.Output != tt.want.callOutput {
				t.Fatal("failed: repository call output is invalid")
			}

			if mockService.CallTimes != tt.want.serviceCallTime {
				t.Fatal("failed: service call times is invalid")
			}
			if mockService.Input != tt.want.serviceInput {
				t.Fatal("failed: service call input is invalid")
			}
			if err := assertBannerConditionServiceOutput(mockService.Output, tt.want.serviceOutput); err != nil {
				t.Fatal(err.Error())
			}

			if err := assertGetShowBannerPageDataOutput(output, tt.want.output); err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}

func assertGetShowBannerPageDataOutput(got *GetShowBannerPageDataOutput, want GetShowBannerPageDataOutput) error {
	if got.PageData != nil {
		if got.PageData.Title() != want.PageData.Title() ||
			got.PageData.Header() != want.PageData.Header() ||
			got.Template != want.Template ||
			got.Error != want.Error {

			gotValues := fmt.Sprintf(
				"title: %v, header: %v, template: %v, error: %v\n",
				got.PageData.Title(),
				got.PageData.Header(),
				got.Template,
				got.Error,
			)
			wantValues := fmt.Sprintf(
				"title: %v, header: %v, template: %v, error: %v\n",
				want.PageData.Title(),
				want.PageData.Header(),
				want.Template,
				want.Error,
			)

			errMessage := fmt.Sprintf("failed: result is invalid\n  Got : %v \n  Want: %v", gotValues, wantValues)
			return errors.New(errMessage)
		}
	} else {
		if got.Template != want.Template ||
			got.Error != want.Error {

			gotValues := fmt.Sprintf(
				"pagedata: %v, template: %v, error: %v\n",
				got.PageData,
				got.Template,
				got.Error,
			)
			wantValues := fmt.Sprintf(
				"pagedata: %v, template: %v, error: %v\n",
				want.PageData,
				want.Template,
				want.Error,
			)

			errMessage := fmt.Sprintf("failed: result is invalid\n  Got : %v \n  Want: %v", gotValues, wantValues)
			return errors.New(errMessage)
		}
	}

	return nil
}

func assertBannerConditionServiceOutput(got *service.BannerConditionOutput, want service.BannerConditionOutput) error {
	if got.Result != want.Result ||
		got.Err != want.Err {
		gotValues := fmt.Sprintf(
			"result: %v, error: %v\n",
			got.Result,
			got.Err,
		)
		wantValues := fmt.Sprintf(
			"result: %v, error: %v\n",
			want.Result,
			want.Err,
		)

		errMessage := fmt.Sprintf("failed: service call output is invalid\n  Got : %v \n  Want: %v", gotValues, wantValues)
		return errors.New(errMessage)
	}

	return nil
}
