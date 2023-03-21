package usecase

import (
	"errors"
	"go-http-server/infrastructure/domain_impl/model"
	domain_model "go-http-server/internal/domain/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockIndexPageRepository struct {
	callTimes               int
	lastInputIndexPage      domain_model.IndexPage
	lastInputResponseWriter http.ResponseWriter
	err                     error
}

func (m *mockIndexPageRepository) SetView(indexPage domain_model.IndexPage, responseWriter http.ResponseWriter) error {
	m.callTimes++
	m.lastInputIndexPage = indexPage
	m.lastInputResponseWriter = responseWriter
	return m.err
}

func Test_GenerateIndexPage_Execute(t *testing.T) {
	var (
		dummyTitle          = "dummy title"
		dummyHeader         = "dummy header"
		dummyResponseWriter = httptest.NewRecorder()
		dummyError          = errors.New("dummy error")

		dummyIndexPage = model.NewIndexPage(
			dummyTitle,
			dummyHeader,
		)
	)

	type (
		input struct {
			title          string
			header         string
			responseWriter *httptest.ResponseRecorder
		}
		mock struct {
			repoErr error
		}
		want struct {
			callTime                int
			callInputIndexPage      domain_model.IndexPage
			callInputResponseWriter http.ResponseWriter
		}
	)

	setup := func(input input, mock mock) (GenerateIndexPageInput, *mockIndexPageRepository) {
		indexPage := model.NewIndexPage(input.title, input.header)
		usecaseInput := GenerateIndexPageInput{
			IndexPage:      indexPage,
			ResponseWriter: input.responseWriter,
		}

		mockIndexPageRepository := &mockIndexPageRepository{
			err: mock.repoErr,
		}

		return usecaseInput, mockIndexPageRepository
	}

	tests := []struct {
		name    string
		input   input
		mock    mock
		want    want
		wantErr error
	}{
		{
			name: "success: no error",
			input: input{
				title:          dummyTitle,
				header:         dummyHeader,
				responseWriter: dummyResponseWriter,
			},
			mock: mock{
				repoErr: nil,
			},
			want: want{
				callTime:                1,
				callInputIndexPage:      dummyIndexPage,
				callInputResponseWriter: dummyResponseWriter,
			},
			wantErr: nil,
		},
		{
			name: "failure: repository error and return error",
			input: input{
				title:          dummyTitle,
				header:         dummyHeader,
				responseWriter: dummyResponseWriter,
			},
			mock: mock{
				repoErr: dummyError,
			},
			want: want{
				callTime:                1,
				callInputIndexPage:      dummyIndexPage,
				callInputResponseWriter: dummyResponseWriter,
			},
			wantErr: dummyError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecaseInput, mockIndexPageRepository := setup(tt.input, tt.mock)
			usecase := NewGenerateIndexPage(mockIndexPageRepository)

			output := usecase.Execute(usecaseInput)
			if mockIndexPageRepository.callTimes != tt.want.callTime {
				t.Fatal("failed: repository call times is invalid")
			}

			if mockIndexPageRepository.lastInputIndexPage.Title() != tt.want.callInputIndexPage.Title() ||
				mockIndexPageRepository.lastInputIndexPage.Header() != tt.want.callInputIndexPage.Header() ||
				mockIndexPageRepository.lastInputResponseWriter != tt.want.callInputResponseWriter {
				t.Fatal("failed: repository call input is invalid")
			}
			if output.Error != tt.wantErr {
				t.Fatal("failed: want error is invalid")
			}
		})
	}
}
