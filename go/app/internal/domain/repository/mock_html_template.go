package repository

import (
	"go-http-server/internal/domain/model"
)

type MockHtmlTemplateRepository struct {
	CallTimes int
	Input     string
	Output    model.HtmlTemplate
	Err       error
}

func NewMockHtmlTemplateRepository() *MockHtmlTemplateRepository {
	return &MockHtmlTemplateRepository{}
}

func (m *MockHtmlTemplateRepository) GetByName(key string) (model.HtmlTemplate, error) {
	m.CallTimes++
	m.Input = key
	return m.Output, m.Err
}
func (m *MockHtmlTemplateRepository) EXPECT(expectedOutput model.HtmlTemplate, expectedErr error) {
	m.Output = expectedOutput
	m.Err = expectedErr
}
