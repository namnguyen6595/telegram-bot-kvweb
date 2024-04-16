package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

func SendResponseImageToChat(chatId int64, imgUrl string) error {
	apiUrl := "https://api.telegram.org/bot" + "6673474158:AAGWhE67vXABkSyL9H-ZCREhSzLrCfvDX48" + "/sendPhoto"
	//"https://api.telegram.org/bot6673474158:AAGWhE67vXABkSyL9H-ZCREhSzLrCfvDX48/sendPhoto"
	commaIndex := strings.Index(imgUrl, ",")
	if commaIndex == -1 {
		fmt.Printf("Error when decode uri")
		return fmt.Errorf("invalid Data URI")
	}
	base64Data := imgUrl[commaIndex+1:]
	photoData, err := base64.StdEncoding.DecodeString(base64Data)
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
