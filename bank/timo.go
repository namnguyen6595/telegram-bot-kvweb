package bank

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"kvweb-bot/helpers"
	"log"
	"net/http"
	"time"
)

type TimoBank struct {
}

type ResponseData struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
	Data    struct {
		Total           int                    `json:"total"`
		TotalBalance    int                    `json:"totalBalance"`
		WorkingBalance  int                    `json:"workingBalance"`
		AccountBalance  int                    `json:"accountBalance"`
		AvailableAmount int                    `json:"availableAmount"`
		CreditLimit     int                    `json:"creditLimit"`
		TxnDateIndex    string                 `json:"txnDateIndex"`
		XidIndex        string                 `json:"xidIndex"`
		PreTxnDateIdx   string                 `json:"preTxnDateIdx"`
		PreXidIdx       string                 `json:"preXidIdx"`
		Products        []interface{}          `json:"products"`
		Format          string                 `json:"format"`
		Items           []*TransactionResponse `json:"items"`
	} `json:"data"`
}

type LogInResponse struct {
	Data AuthorizeData `json:"data"`
	//SecData struct {
	//	ExpiryTime string `json:"expiryTime"`
	//	TimeZone   string `json:"timeZone"`
	//	ServerTime string `json:"serverTime"`
	//	RfToken    string `json:"rfToken"`
	//	Token      string `json:"token"`
	//} `json:"secData"`
}

type AuthorizeData struct {
	UserID int    `json:"userId"`
	Lang   string `json:"lang"`
	Token  string `json:"token"`
}

func (t *TimoBank) GetTransaction() ([]*TransactionResponse, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	env, envErr := godotenv.Read(".env")
	if envErr != nil {
		log.Fatalf("Error when read file env: %v", envErr)
	}
	url := "https://app2.timo.vn/user/account/transaction/list"
	client := http.Client{
		Timeout: time.Second * 120, // Timeout after 2 seconds
	}
	now := time.Now()
	fromDate, toDate := helpers.GetFirstAndLastDayOfMonth(now, "02/01/2006")
	timoBody := map[string]any{
		"format":      "group",
		"index":       0,
		"offset":      -1,
		"accountNo":   "9021543768115",
		"accountType": "1025",
		"fromDate":    "01/01/2015",
		"toDate":      "15/04/2024",
		"filter": map[string]string{
			"moneyType": "all",
			"fromDate":  fromDate,
			"toDate":    toDate,
		},
	}
	token := env["BANK_TOKEN"]
	bodyReq, _ := json.Marshal(timoBody)
	// Handle token unauthorization
	request, err := http.NewRequest(http.MethodGet, "https://app2.timo.vn/login/quickCode/check", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("token", token)
	resRaw, err := client.Do(request)
	//if resRaw.StatusCode == http.StatusUnauthorized {
	//	bodyReq, _ = json.Marshal(map[string]string{
	//		"username": "0392126595",
	//		"password": "485620e6d3bad737c32d4d149347d6425c2a523cc5b59a87d00d319eceeecaf65eadfef8007165c0976e6824dd04d75161f63a3a142f95f85669d942969091a8",
	//		"lang":     "vn",
	//	})
	//	request, err = http.NewRequest(http.MethodPost, "https://app2.timo.vn/login", bytes.NewReader(bodyReq))
	//	request.Header.Set("Content-Type", "application/json")
	//	authorizeData, err := client.Do(request)
	//	var loginResponse *LogInResponse
	//	if err != nil {
	//		log.Printf("Error when re-authorize: %v", err)
	//		return nil, err
	//	}
	//	bodyAuthorize, readErr := ioutil.ReadAll(authorizeData.Body)
	//	if readErr != nil {
	//		log.Printf("err: %v", readErr)
	//		return nil, readErr
	//	}
	//
	//	err = json.Unmarshal(bodyAuthorize, &loginResponse)
	//	if err != nil {
	//		log.Printf("Error when parse data login: %v", err)
	//		return nil, err
	//	}
	//
	//	//token = loginResponse.Data.Token
	//}
	request, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyReq))

	if err != nil {
		log.Printf("Error when create new request: %v", err)
	}
	request.Header.Set("token", token)
	bankResponse := &ResponseData{}

	resRaw, err = client.Do(request)

	if err != nil {
		log.Printf("Error when send requst. %v", err)
		return nil, err
	}

	if resRaw.StatusCode != http.StatusOK {
		log.Printf("error when request: %v", map[string]interface{}{
			"code":  resRaw.StatusCode,
			"error": err,
			"token": token,
		})
		return nil, err
	}

	body, readErr := ioutil.ReadAll(resRaw.Body)
	if readErr != nil {
		log.Printf("err: %v", readErr)
		return nil, readErr
	}

	err = json.Unmarshal(body, &bankResponse)
	if err != nil {
		log.Printf("err: %v", err)
		return nil, err
	}

	return bankResponse.Data.Items, nil
}
