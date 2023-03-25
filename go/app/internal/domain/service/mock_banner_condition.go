package service

type MockBannerConditionService struct {
	CallTimes int
	Input     BannerConditionInput
	Output    *BannerConditionOutput
}

func NewMockBannerConditionService() *MockBannerConditionService {
	return &MockBannerConditionService{}
}

func (m *MockBannerConditionService) IsBannerDisplayed(bannerConditionInput BannerConditionInput) *BannerConditionOutput {
	m.CallTimes++
	m.Input = bannerConditionInput
	return m.Output
}
func (m *MockBannerConditionService) EXPECT(expectedOutput *BannerConditionOutput) {
	m.Output = expectedOutput
}
