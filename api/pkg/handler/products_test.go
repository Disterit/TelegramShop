package handler

import (
	Telegram_Market "Telegram-Market"
	mock_service "Telegram-Market/api/pkg/service/mocks"
	"testing"
)

func TestHandler_CreateProduct(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProducts, products Telegram_Market.Products)

	testTable := []struct {
		name                string
		userInput           string
		bodyInput           string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "ok",
		},
	}
}
