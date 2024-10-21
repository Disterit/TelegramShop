package handler

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/handler/kafkaHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const er = "error: "

func (h *Handler) CreateUser(c *gin.Context) {
	const op = "api.pkg.handler.user.CreateUser()"

	userIdStr := c.Param("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.User.Create(userId)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "user successfully created", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) GetUserById(c *gin.Context) {
	const op = "api.pkg.handler.user.GetUserById()"

	userIdStr := c.Param("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	res, err := h.service.User.GetUserById(userId)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "user successfully retrieved", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	const op = "api.pkg.handler.user.GetAllUsers()"

	res, err := h.service.User.GetUsers()
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "users successfully retrieved", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	const op = "api.pkg.handler.user.UpdateUser()"

	userIdStr := c.Param("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input Telegram_Market.Users
	if err := c.BindJSON(&input); err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.User.UpdateUser(userId, input)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "user successfully updated", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	const op = "api.pkg.handler.user.DeleteUser()"

	userIdStr := c.Param("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.User.DeleteUser(userId)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, er+err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "user successfully deleted", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
