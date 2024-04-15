package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	qr_code "kvweb-bot/qr-code"
	"log"
	"mime/multipart"
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

	messageData, err := json.MarshalIndent(update, "", "    ") // Sử dụng MarshalIndent để định dạng JSON cho dễ đọc
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}

	log.Printf("Receive message from data: %v", string(messageData))

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

func sendResponseToChat(chatId int64, imgUrl string) error {
	apiUrl := "https://api.telegram.org/bot" + "6673474158:AAGWhE67vXABkSyL9H-ZCREhSzLrCfvDX48" + "/sendPhoto"
	//"https://api.telegram.org/bot6673474158:AAGWhE67vXABkSyL9H-ZCREhSzLrCfvDX48/sendPhoto"
	photoData, err := base64.StdEncoding.DecodeString(imgUrl)
	if err != nil {
		log.Printf("Error when decode image: %v", err)
		return err
	}

	// Tạo một yêu cầu HTTP POST multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("photo", "image.jpg")
	if err != nil {
		log.Printf("Error when create image: %v", err)
		return err
	}

	if _, err = io.Copy(part, bytes.NewReader(photoData)); err != nil {
		log.Printf("Error when copy image: %v", err)
		return err
	}

	_ = writer.WriteField("chat_id", fmt.Sprintf("%d", chatId))
	writer.Close()

	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		log.Printf("Erro when create request. %v", err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Gửi yêu cầu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error when request. %v", err)
		return err
	}
	defer resp.Body.Close()

	// Xử lý phản hồi nếu cần
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response from Telegram: %v", resp)
	}
	return nil
}
