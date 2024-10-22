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

func TestHandler_CreateProduct(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProducts, products Telegram_Market.Products)

	testTable := []struct {
		name                string
		inputUser           Telegram_Market.Products
		inputBody           string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "ok",
			inputUser: Telegram_Market.Products{
				Description:         "Sample product description",
				PhotoUrl:            "https://example.com/product-image.jpg",
				Price:               29.99,
				Quantity:            100,
				LocationCity:        "Moscow",
				LocationCoordinates: "55.7558, 37.6173",
				HiddenPhotoUrl:      "https://example.com/hidden-product-image.jpg",
				HiddenDescription:   "This is hidden product description",
				PaidFlag:            true,
			},
			inputBody: `{"description": "Sample product description", 
						"photo_url": "https://example.com/product-image.jpg", 
						"price": 29.99, 
						"quantity": 100, "location_city": 
						"Moscow", "location_coordinates": "55.7558, 37.6173", 
						"hidden_photo_url": "https://example.com/hidden-product-image.jpg", 
						"hidden_description": "This is hidden product description", 
						"paid_flag": true}`,
			mockBehavior: func(s *mock_service.MockProducts, products Telegram_Market.Products) {
				s.EXPECT().CreateProduct(products).Return(int64(1), nil)
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

			createProduct := mock_service.NewMockProducts(c)
			testCase.mockBehavior(createProduct, testCase.inputUser)

			services := &service.Service{Products: createProduct}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.POST("/products", handler.CreateProduct)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/products", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_GetProductById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProducts, productId int64)

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
			mockBehavior: func(s *mock_service.MockProducts, productId int64) {
				s.EXPECT().GetProductById(productId).Return(Telegram_Market.Products{
					Id:                  int64(1),
					Description:         "Sample product description",
					PhotoUrl:            "https://example.com/product-image.jpg",
					Price:               29.99,
					Quantity:            100,
					LocationCity:        "Moscow",
					LocationCoordinates: "55.7558, 37.6173",
					HiddenPhotoUrl:      "https://example.com/hidden-product-image.jpg",
					HiddenDescription:   "This is hidden product description",
					PaidFlag:            true,
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"description":"Sample product description","photo_url":"https://example.com/product-image.jpg","price":29.99,"quantity":100,"location_city":"Moscow","location_coordinates":"55.7558, 37.6173","hidden_photo_url":"https://example.com/hidden-product-image.jpg","hidden_description":"This is hidden product description","paid_flag":true}`,
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

			productById := mock_service.NewMockProducts(c)
			testCase.mockBehavior(productById, testCase.userInput)

			services := &service.Service{Products: productById}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.GET("/products/:id", handler.GetProductById)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/products/1", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_UpdateProduct(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProducts, productsId int64, products Telegram_Market.UpdateProducts)

	testTable := []struct {
		name                string
		inputUser           int64
		inputBody           Telegram_Market.UpdateProducts
		inputBodyJson       string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "ok",
			inputUser: int64(1),
			inputBody: Telegram_Market.UpdateProducts{
				Price: func() *float64 {
					val := 150.0
					return &val
				}(),
				LocationCity: func() *string {
					val := "aboba"
					return &val
				}(),
			},
			inputBodyJson: `{"price":150.0,"location_city":"aboba"}`,
			mockBehavior: func(s *mock_service.MockProducts, productsId int64, products Telegram_Market.UpdateProducts) {
				s.EXPECT().UpdateProduct(productsId, products).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"status":"ok"}`,
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

			updateLocation := mock_service.NewMockProducts(c)
			testCase.mockBehavior(updateLocation, testCase.inputUser, testCase.inputBody)

			services := &service.Service{Products: updateLocation}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.PUT("/products/:id", handler.UpdateProduct)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewBufferString(testCase.inputBodyJson))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteProduct(t *testing.T) {
	type mockBehavior func(s *mock_service.MockProducts, productsId int64)

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
			mockBehavior: func(s *mock_service.MockProducts, productsId int64) {
				s.EXPECT().DeleteProduct(productsId).Return(nil)
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

			deleteProduct := mock_service.NewMockProducts(c)
			testCase.mockBehavior(deleteProduct, testCase.userInput)

			services := &service.Service{Products: deleteProduct}
			handler := NewHandler(services, writer)

			r := gin.New()
			r.DELETE("/products/:id", handler.DeleteProduct)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
