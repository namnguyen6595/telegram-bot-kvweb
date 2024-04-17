package qr_code

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type GenerateQrRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

type QrResponse struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
	Data struct {
		AcpID       int    `json:"acpId"`
		AccountName string `json:"accountName"`
		QrCode      string `json:"qrCode"`
		QrDataURL   string `json:"qrDataURL"`
	} `json:"data"`
}

type QrData struct {
	QrCode    string `json:"qrCode"`
	QrDataURL string `json:"qrDataURL"`
}

func GenerateVietQrCode(req *GenerateQrRequest) (*QrData, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	env, _ := godotenv.Read(".env")
	clientKey := env["VIETQR_CLIENT_KEY"]
	apiKey := env["VIETQR_API_KEY"]
	url := fmt.Sprintf("https://api.vietqr.io/v2/generate")
	client := http.Client{
		Timeout: time.Second * 120, // Timeout after 2 seconds
	}
	vietQrBody := map[string]string{
		"accountNo":   "9021543768115",
		"accountName": "NGUYEN THANH NAM",
		"acqId":       "963388",
		"amount":      fmt.Sprintf("%v", req.Amount*1000),
		"addInfo":     req.Description,
		"format":      "text",
		"template":    "print",
	}
	bodyReq, _ := json.Marshal(vietQrBody)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyReq))

	if err != nil {
		log.Fatalf("Error when create request. %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-client-id", clientKey)
	request.Header.Set("x-api-key", apiKey)
	response := &QrResponse{}
	resRaw, err := client.Do(request)

	if err != nil {
		log.Fatalf("Error when send requst. %v", err)
	}

	body, readErr := ioutil.ReadAll(resRaw.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	return &QrData{
		QrCode:    response.Data.QrCode,
		QrDataURL: response.Data.QrDataURL,
	}, nil
}
