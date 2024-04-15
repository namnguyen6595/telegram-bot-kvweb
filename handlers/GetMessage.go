package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	qr_code "kvweb-bot/qr-code"
	"log"
	"net/http"
	"net/url"
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

	messageData, err := json.MarshalIndent(update, "", "    ") // Sử dụng MarshalIndent để định dạng JSON cho dễ đọc
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	log.Printf("Receive message from data: %v", messageData)

	// Kiểm tra xem tin nhắn có phải là lệnh /topup
	if update.Message.Text == "/topup" {
		log.Printf("Data message: %v", update)
		qrCode, _ := qr_code.GenerateVietQrCode(&qr_code.GenerateQrRequest{
			Amount: 200,
		})
		sendResponseToChat(update.Message.Chat.ID, qrCode.QrDataURL)
		ctx.JSON(http.StatusOK, qrCode)
		return
	}

	ctx.JSONP(http.StatusOK, map[string]string{
		"data": "update",
	})
}

func sendResponseToChat(chatId int64, imgUrl string) {
	apiUrl := "https://api.telegram.org/bot6673474158:AAGWhE67vXABkSyL9H-ZCREhSzLrCfvDX48/sendPhoto"
	data := url.Values{}
	data.Set("chat_id", fmt.Sprintf("%d", chatId))
	data.Set("photo", imgUrl)

	resp, err := http.PostForm(apiUrl, data)
	if err != nil {
		log.Fatalf("Error sending photo: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Received non-OK response from Telegram: %s", resp.Status)
	}
}
