package handler

import (
	Telegram_Market "Telegram-Market"
	"Telegram-Market/api/pkg/handler/kafkaHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateWallet(c *gin.Context) {
	const op = "api.pkg.handler.wallets.CreateWallet()"
	var wallet Telegram_Market.Wallet

	if err := c.BindJSON(&wallet); err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Wallets.CreateWallet(wallet)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully created", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetWalletById(c *gin.Context) {
	const op = "api.pkg.handler.wallets.GetWalletById()"

	id := c.Param("id")
	WalletID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Wallets.GetWalletById(WalletID)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully retrieved", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllWallets(c *gin.Context) {
	const op = "api.pkg.handler.wallets.GetAllWallets()"

	res, err := h.service.Wallets.GetAllWallets()
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "products successfully retrieved", op)
	c.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateWallet(c *gin.Context) {
	const op = "api.pkg.handler.wallets.UpdateWallet()"

	id := c.Param("id")
	WalletID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var wallet Telegram_Market.Wallet
	if err := c.BindJSON(&wallet); err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Wallets.UpdateWallet(WalletID, wallet)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully updated", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}

func (h *Handler) DeleteWallet(c *gin.Context) {
	const op = "api.pkg.handler.wallets.DeleteWallet()"

	id := c.Param("id")
	WalletID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteWallet(WalletID)
	if err != nil {
		kafkaHandler.KafkaResponse(h.writer, err.Error(), op)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	kafkaHandler.KafkaResponse(h.writer, "product successfully deleted", op)
	c.JSON(http.StatusOK, map[string]interface{}{
		"ok": "ok",
	})
}
