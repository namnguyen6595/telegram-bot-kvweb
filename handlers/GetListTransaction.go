package handlers

import (
	"github.com/gin-gonic/gin"
	"kvweb-bot/models"
)

type GetListTransaction struct {
	BalanceDmsClient *models.BalanceDmsClient
}

func (h *GetListTransaction) NewServe(ctx *gin.Context) {

}
