package handler

import (
	"Telegram-Market/api/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Handler struct {
	service *service.Service
	writer  *kafka.Writer
}

func NewHandler(service *service.Service, writer *kafka.Writer) *Handler {
	return &Handler{service: service, writer: writer}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/:id", h.CreateUser)   // добавить пользователя
		user.GET("/:id", h.GetUserById)   // получить информацию по 1 пользователю
		user.GET("/", h.GetAllUsers)      // получить информацию по всем пользователям
		user.PUT("/:id", h.UpdateUser)    // изменить баланс ползователю по id
		user.DELETE("/:id", h.DeleteUser) // удалить пользователя по id
	}

	products := router.Group("/products")
	{
		products.POST("/", h.CreateProduct)      // добавить продукт
		products.GET("/:id", h.GetProductById)   // получить информацию по 1 продукту
		products.GET("/", h.GetAllProducts)      // получить информацию по всем продуктам
		products.PUT("/:id", h.UpdateProduct)    // изменить информацию о продукте по id
		products.DELETE("/:id", h.DeleteProduct) // удалить продукт по id
	}

	locations := router.Group("/locations")
	{
		locations.POST("/", h.CreateLocation)      // добавить локацию
		locations.GET("/:id", h.GetLocationById)   // получить информацию по 1 локации
		locations.GET("/", h.GetAllLocations)      // получить информацию по всем локациям
		locations.PUT("/:id", h.UpdateLocation)    // изменить информацию о локацкии по id
		locations.DELETE("/:id", h.DeleteLocation) // удалить локацию по id
	}

	wallets := router.Group("/wallets")
	{
		wallets.POST("/", h.CreateWallet)      // добавить кошелёк
		wallets.GET("/:id", h.GetWalletById)   // получить информацию по 1 кошельку
		wallets.GET("/", h.GetAllWallets)      // получить информацию по всем кошелькам
		wallets.PUT("/:id", h.UpdateWallet)    // изменить информацию о кошельке по id
		wallets.DELETE("/:id", h.DeleteWallet) // удалить кошелёк по id
	}

	return router
}
