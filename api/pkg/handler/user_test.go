package handler

import (
	Telegram_Market "Telegram-Market"
	service "Telegram-Market/api/pkg/service"
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

func TestHandler_CreateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUser, userId int64)

	testTable := []struct {
		name                string
		inputUser           int64
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputUser: int64(421352132),
			mockBehavior: func(s *mock_service.MockUser, userId int64) {
				s.EXPECT().CreateUser(userId).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"ok":"ok"}`,
		},
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "TelegramShop",
		Balancer: &kafka.LeastBytes{},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			create := mock_service.NewMockUser(c)
			testCase.mockBehavior(create, testCase.inputUser)

			services := &service.Service{User: create}
			handler := NewHandler(services, writer)
			//Test server
			r := gin.New()
			r.POST("/user/:id", handler.CreateUser)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user/421352132", nil)

			//perform request
			r.ServeHTTP(w, req)

			//assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetUserById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUser, userId int64)

	testTable := []struct {
		name                string
		inputUser           int64
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputUser: int64(142135213),
			mockBehavior: func(s *mock_service.MockUser, userId int64) {
				s.EXPECT().GetUserById(userId).Return(Telegram_Market.Users{
					Id:      1,
					UserId:  142135213,
					Balance: 0},
					nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"user_id":142135213,"balance":0}`,
		},
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "TelegramShop",
		Balancer: &kafka.LeastBytes{},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			getUser := mock_service.NewMockUser(c)
			testCase.mockBehavior(getUser, testCase.inputUser)

			services := &service.Service{User: getUser}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.GET("/user/:id", handler.GetUserById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/user/142135213", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_UpdateUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUser, userId int64, user Telegram_Market.Users)

	tableTest := []struct {
		name                string
		userInput           int64
		bodyInput           Telegram_Market.Users
		bodyInputJson       string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:          "ok",
			userInput:     int64(124213412),
			bodyInput:     Telegram_Market.Users{Balance: 150.0},
			bodyInputJson: `{"balance": 150}`,
			mockBehavior: func(s *mock_service.MockUser, userId int64, user Telegram_Market.Users) {
				s.EXPECT().UpdateUser(userId, user).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"ok":"ok"}`,
		},
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "TelegramShop",
		Balancer: &kafka.LeastBytes{},
	})

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			updateUser := mock_service.NewMockUser(c)
			testCase.mockBehavior(updateUser, testCase.userInput, testCase.bodyInput)

			services := &service.Service{User: updateUser}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.PUT("/user/:id", handler.UpdateUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/user/124213412", bytes.NewBufferString(testCase.bodyInputJson))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUser, userId int64)

	testTable := []struct {
		name                string
		userInput           int64
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			userInput: int64(54213241),
			mockBehavior: func(s *mock_service.MockUser, userId int64) {
				s.EXPECT().DeleteUser(userId).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"ok":"ok"}`,
		},
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "TelegramShop",
		Balancer: &kafka.LeastBytes{},
	})

	for _, testCase := range testTable {
		c := gomock.NewController(t)
		defer c.Finish()

		updateUser := mock_service.NewMockUser(c)
		testCase.mockBehavior(updateUser, testCase.userInput)

		services := &service.Service{User: updateUser}
		handler := NewHandler(services, writer)

		r := gin.New()
		r.DELETE("/user/:id", handler.DeleteUser)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/user/54213241", nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, testCase.expectedStatusCode, w.Code)
		assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
	}
}
