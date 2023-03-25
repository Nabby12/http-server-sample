package usecase

import (
	"errors"
	"fmt"
	inframodel "go-http-server/infrastructure/domain_impl/model"
	"go-http-server/internal/domain/model"
	"go-http-server/internal/domain/repository"
	"testing"
)

func Test_GetIndexPageData_Execute(t *testing.T) {
	var (
		dummyTitle            = "dummy title"
		dummyHeader           = "dummy header"
		dummyKey              = IndexKey
		dummyTemplateFileName = "dummy html"
		dummyPublicPath       = "dummy public"
		dummyTemplateFullPath = "dummy public/dummy html"
		dummyError            = errors.New("dummy error")

		dummyHtmlPage = inframodel.NewHtmlTemplate(
			dummyTemplateFileName,
			dummyPublicPath,
		)
		dummyIndexPage = inframodel.NewIndexPage(
			dummyTitle,
			dummyHeader,
		)
	)

	type (
		input struct {
			title     string
			header    string
			repoInput string
		}
		mock struct {
			repoOutput model.HtmlTemplate
			repoErr    error
		}
		want struct {
			callTime     int
			callInputKey string
			callOutput   model.HtmlTemplate
			output       GetIndexPageDataOutput
		}
	)

	setup := func(input input, mock mock) (*getIndexPageDataImpl, *repository.MockHtmlTemplateRepository) {
		mockHtmlTemplateRepository := repository.NewMockHtmlTemplateRepository()
		mockHtmlTemplateRepository.EXPECT(
			mock.repoOutput,
			mock.repoErr,
		)

		usecase := &getIndexPageDataImpl{
			mockHtmlTemplateRepository,
		}
		return usecase, mockHtmlTemplateRepository
	}

	tests := []struct {
		name  string
		input input
		mock  mock
		want  want
	}{
		{
			name: "success: no error",
			input: input{
				title:     dummyTitle,
				header:    dummyHeader,
				repoInput: dummyKey,
			},
			mock: mock{
				repoOutput: dummyHtmlPage,
				repoErr:    nil,
			},
			want: want{
				callTime:     1,
				callInputKey: dummyKey,
				callOutput:   dummyHtmlPage,
				output: GetIndexPageDataOutput{
					PageData: dummyIndexPage,
					Template: dummyTemplateFullPath,
					Error:    nil,
				},
			},
		},
		{
			name: "failure: repository error and return error",
			input: input{
				title:     dummyTitle,
				header:    dummyHeader,
				repoInput: dummyKey,
			},
			mock: mock{
				repoOutput: nil,
				repoErr:    dummyError,
			},
			want: want{
				callTime:     1,
				callInputKey: dummyKey,
				callOutput:   nil,
				output: GetIndexPageDataOutput{
					PageData: nil,
					Template: "",
					Error:    dummyError,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase, mockRepo := setup(tt.input, tt.mock)

			output := usecase.Execute(GetIndexPageDataInput{
				Title:  tt.input.title,
				Header: tt.input.header,
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

			if err := assertOutput(output, tt.want.output); err != nil {
				t.Fatal(err.Error())
			}
		})
	}
}

func assertOutput(got *GetIndexPageDataOutput, want GetIndexPageDataOutput) error {
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
