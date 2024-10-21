package handler

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/handler/kafkaHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	const op = "api.pkg.handler.products.CreateProduct()"

	var input Telegram_Market.Products
	if err := c.BindJSON(&input); err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Products.CreateProduct(input)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully retrieved", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetProductById(c *gin.Context) {
	const op = "api.pkg.handler.products.GetProductById()"

	productIdStr := c.Param("id")
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	res, err := h.service.Products.GetProductById(productId)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully retrieved", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	const op = "api.pkg.handler.products.GetAllProducts()"

	res, err := h.service.Products.GetAllProducts()
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "products successfully getting", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	const op = "api.pkg.handler.products.UpdateProduct()"

	productIdStr := c.Param("id")
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input Telegram_Market.UpdateProducts
	if err := c.BindJSON(&input); err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Products.UpdateProduct(productId, input)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully updated", op)
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	const op = "api.pkg.handler.products.DeleteProduct()"

	userIdStr := c.Param("id")
	productId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.Products.DeleteProduct(productId)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully deleted", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
