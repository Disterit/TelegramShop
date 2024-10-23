package handler

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/service"
	mock_service "Telegram-Market/api/pkg/service/mocks"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
	"net/http"
	"net/http/httptest"
	"testing"
)

var writeTest *kafka.Writer

func init() {
	writeTest = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "TelegramShop",
		Balancer: &kafka.LeastBytes{},
	})
}

func TestHandler_CreateWallet(t *testing.T) {
	type mockBehavior func(s *mock_service.MockWallets, wallet Telegram_Market.Wallet)

	testTable := []struct {
		name                string
		inputUser           Telegram_Market.Wallet
		inputBody           string
		mockBehavior        mockBehavior
		exceptedStatusCode  int
		exceptedRequestBody string
	}{
		{
			name: "ok",
			inputUser: Telegram_Market.Wallet{
				ID:   int64(1),
				Name: "Monero",
				Addresses: []string{
					"84gjsd74h3gkjsdhg74kdhsk73j",
					"akfh874gfhshfjkdhskd7393kd",
					"hfkdj784kfjhsdjd39fjhfkdj"},
			},
			inputBody: `{"id": 1, "name": "Monero", "addresses": [ "84gjsd74h3gkjsdhg74kdhsk73j", "akfh874gfhshfjkdhskd7393kd", "hfkdj784kfjhsdjd39fjhfkdj"]}`,
			mockBehavior: func(s *mock_service.MockWallets, wallet Telegram_Market.Wallet) {
				s.EXPECT().CreateWallet(wallet).Return(int64(1), nil)
			},
			exceptedStatusCode:  200,
			exceptedRequestBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			createWallet := mock_service.NewMockWallets(c)
			testCase.mockBehavior(createWallet, testCase.inputUser)

			services := &service.Service{Wallets: createWallet}
			handler := NewHandler(services, writeTest)

			r := gin.New()
			r.POST("/wallets", handler.CreateWallet)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/wallets", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.exceptedStatusCode, w.Code)
			assert.Equal(t, testCase.exceptedRequestBody, w.Body.String())
		})
	}
}
