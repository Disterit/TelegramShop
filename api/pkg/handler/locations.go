package handler

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/handler/kafkaHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateLocation(c *gin.Context) {
	const op = "api.pkg.handler.locations.CreateLocation()"

	var input Telegram_Market.Locations
	if err := c.BindJSON(&input); err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Locations.CreateLocation(input)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "locations successfully created", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetLocationById(c *gin.Context) {
	const op = "api.pkg.handler.locations.GetLocationById()"

	LocationIdStr := c.Param("id")
	LocationId, err := strconv.ParseInt(LocationIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Locations.GetLocationById(LocationId)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "location successfully retrieved", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllLocations(c *gin.Context) {
	const op = "api.pkg.handler.locations.GetAllLocations()"

	res, err := h.service.Locations.GetAllLocations()
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "locations successfully retrieved", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateLocation(c *gin.Context) {
	const op = "api.pkg.handler.locations.UpdateLocation()"

	LocationIdStr := c.Param("id")
	LocationId, err := strconv.ParseInt(LocationIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input Telegram_Market.UpdateLocations
	if err := c.BindJSON(&input); err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Locations.UpdateLocations(LocationId, input)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "location successfully updated", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) DeleteLocation(c *gin.Context) {
	const op = "api.pkg.handler.locations.DeleteLocation()"

	LocationIdStr := c.Param("id")
	LocationId, err := strconv.ParseInt(LocationIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Locations.DeleteLocation(LocationId)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "location successfully deleted", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
