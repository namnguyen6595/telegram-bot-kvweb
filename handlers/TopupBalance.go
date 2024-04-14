package handlers

import (
	"github.com/gin-gonic/gin"
	"kvweb-bot/models"
	qr_code "kvweb-bot/qr-code"
	"net/http"
)

type TopupBalanceHandler struct {
	BalanceDmsClient *models.BalanceDmsClient
}

func (h *TopupBalanceHandler) NewServe(ctx *gin.Context) {
	data, _ := qr_code.GenerateVietQrCode(&qr_code.GenerateQrRequest{
		Name:        "nam",
		Description: "Nam topup",
		Amount:      200,
	})

	ctx.JSONP(http.StatusOK, data)
}
