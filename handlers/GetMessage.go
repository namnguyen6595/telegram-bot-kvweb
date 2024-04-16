package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"kvweb-bot/bank"
	"kvweb-bot/helpers"
	qr_code "kvweb-bot/qr-code"
	"log"
	"net/http"
)

type GetMessageHandler struct {
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type ListTransactionResponse struct {
	Success      bool                        `json:"success"`
	Transactions []*bank.TransactionResponse `json:"transactions"`
}

type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	From struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
	} `json:"from"`
}

type Chat struct {
	ID int64 `json:"id"`
}

func (h *GetMessageHandler) NewServe(ctx *gin.Context) {
	banks := bank.InitialBanks()
	var update Update
	if err := ctx.BindJSON(&update); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	messageData, err := json.MarshalIndent(update, "", "    ") // Sử dụng MarshalIndent để định dạng JSON cho dễ đọc
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	log.Printf("Receive message from data: %v", string(messageData))

	// handle with message is topup
	if update.Message.Text == "/dongquy" {
		userID := update.Message.From.ID
		firstName := update.Message.From.FirstName
		lastName := update.Message.From.LastName
		username := update.Message.From.Username
		chatID := update.Message.Chat.ID

		// In ra console hoặc xử lý thông tin
		log.Printf("Message from %s %s (Username: %s, UserID: %d) in chat %d", firstName, lastName, username, userID, chatID)
		qrCode, _ := qr_code.GenerateVietQrCode(&qr_code.GenerateQrRequest{
			Amount:      200,
			Description: fmt.Sprintf("%v %v đóng tiền quỹ", userID, firstName+lastName),
			Name:        fmt.Sprintf("%v", firstName+lastName),
		})
		err = helpers.SendResponseImageToChat(update.Message.Chat.ID, qrCode.QrDataURL)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, qrCode)
		return
	}

	// Xem lich su giao dich
	if update.Message.Text == "/lichsu" {
		transactions, err := banks.GetTransaction()
		if err != nil {
			ctx.JSON(http.StatusNotFound, map[string]string{
				"error": err.Error(),
			})
			return
		}
		ctx.JSONP(http.StatusOK, &ListTransactionResponse{
			Success:      true,
			Transactions: transactions,
		})
	}

	ctx.JSONP(http.StatusOK, &ListTransactionResponse{
		Success: true,
	})
}
