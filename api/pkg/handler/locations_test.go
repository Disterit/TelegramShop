package handler

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/service"
	mock_service "Telegram-Market/api/pkg/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateLocation(t *testing.T) {
	type mockBehavior func(s *mock_service.MockLocations, locations Telegram_Market.Locations)

	testTable := []struct {
		name                string
		userInput           string
		bodyInput           Telegram_Market.Locations
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			userInput: `{"country": "Russia", "city": "Moscow", "city_district": "Tverskoy"}`,
			bodyInput: Telegram_Market.Locations{
				Country:       "Russia",
				City:          "Moscow",
				City_district: "Tverskoy",
			},
			mockBehavior: func(s *mock_service.MockLocations, locations Telegram_Market.Locations) {
				s.EXPECT().CreateLocation(locations).Return(int64(1), nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
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

			createLocations := mock_service.NewMockLocations(c)
			testCase.mockBehavior(createLocations, testCase.bodyInput)

			services := &service.Service{Locations: createLocations}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.POST("/locations", handler.CreateLocation)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/locations", bytes.NewBufferString(testCase.userInput))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})

	}
}

func TestHandler_GetLocationById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockLocations, locationId int64)

	testTable := []struct {
		name                string
		userInput           int64
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			userInput: int64(1),
			mockBehavior: func(s *mock_service.MockLocations, locationId int64) {
				s.EXPECT().GetLocationById(locationId).Return(Telegram_Market.Locations{
					Id:            1,
					Country:       "Russia",
					City:          "Moscow",
					City_district: "Tverskoy",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"country":"Russia","city":"Moscow","city_district":"Tverskoy"}`,
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

			getLocationsById := mock_service.NewMockLocations(c)
			testCase.mockBehavior(getLocationsById, testCase.userInput)

			services := &service.Service{Locations: getLocationsById}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.GET("/locations/:id", handler.GetLocationById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/locations/1", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_UpdateLocation(t *testing.T) {
	type mockBehavior func(s *mock_service.MockLocations, locationsId int64, locations Telegram_Market.UpdateLocations)

	testTable := []struct {
		name                string
		userInputId         int64
		userInput           string
		userInputBody       Telegram_Market.UpdateLocations
		userBody            string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "ok",
			userInputId: int64(1),
			userInput:   `{"country": "Russia", "city": "Moscow", "city_district": "Tverskoy"}`,
			userInputBody: Telegram_Market.UpdateLocations{
				Country: func() *string {
					val := "Russia"
					return &val
				}(),
				City: func() *string {
					val := "Moscow"
					return &val
				}(),
				City_district: func() *string {
					val := "Tverskoy"
					return &val
				}(),
			},
			mockBehavior: func(s *mock_service.MockLocations, locationsId int64, locations Telegram_Market.UpdateLocations) {
				s.EXPECT().UpdateLocations(locationsId, locations).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"ok":"ok"}`,
		},
		{
			name:        "service error",
			userInputId: int64(1),
			userInput:   `{"country": "Russia", "city": "Moscow", "city_district": "Tverskoy"}`,
			userInputBody: Telegram_Market.UpdateLocations{
				Country: func() *string {
					val := "Russia"
					return &val
				}(),
				City: func() *string {
					val := "Moscow"
					return &val
				}(),
				City_district: func() *string {
					val := "Tverskoy"
					return &val
				}(),
			},
			mockBehavior: func(s *mock_service.MockLocations, locationsId int64, locations Telegram_Market.UpdateLocations) {
				s.EXPECT().UpdateLocations(locationsId, locations).Return(errors.New("internal error"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"internal error"}`,
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

			updateLocations := mock_service.NewMockLocations(c)
			testCase.mockBehavior(updateLocations, testCase.userInputId, testCase.userInputBody)

			services := &service.Service{Locations: updateLocations}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.PUT("/locations/:id", handler.UpdateLocation)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/locations/1", bytes.NewBufferString(testCase.userInput))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteLocation(t *testing.T) {
	type mockBehavior func(s *mock_service.MockLocations, locationsId int64)

	testTable := []struct {
		name                string
		userInput           int64
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			userInput: int64(1),
			mockBehavior: func(s *mock_service.MockLocations, locationsId int64) {
				s.EXPECT().DeleteLocation(locationsId).Return(nil)
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
			c := gomock.NewController(t)
			defer c.Finish()

			deleteLocation := mock_service.NewMockLocations(c)
			testCase.mockBehavior(deleteLocation, testCase.userInput)

			services := &service.Service{Locations: deleteLocation}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.DELETE("/locations/:id", handler.DeleteLocation)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/locations/1", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
