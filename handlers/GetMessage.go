package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GetMessageHandler struct {
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID int64 `json:"id"`
}

func (h *GetMessageHandler) NewServe(ctx *gin.Context) {
	var update Update
	if err := ctx.BindJSON(&update); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Kiểm tra xem tin nhắn có phải là lệnh /topup
	if update.Message.Text == "/topup" {

		ctx.JSON(http.StatusOK, update)
		return
	}

	ctx.JSONP(http.StatusOK, map[string]string{
		"data": "update",
	})
}
