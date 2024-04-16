package handlers

import (
	"github.com/gin-gonic/gin"
	"kvweb-bot/bank"
	"log"
	"net/http"
)

type GetListTransaction struct {
}

func (h *GetListTransaction) NewServe(ctx *gin.Context) {
	banks := bank.InitialBanks()

	response, err := banks.GetTransaction()

	if err != nil {
		log.Printf("Error: %v", err)
		ctx.JSON(http.StatusBadGateway, map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success":      true,
		"transactions": response,
	})
}
