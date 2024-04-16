package bank

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func (t *TimoBank) GetTransaction() ([]*TransactionResponse, error) {
	url := "https://app2.timo.vn/user/account/transaction/list"
	client := http.Client{
		Timeout: time.Second * 120, // Timeout after 2 seconds
	}
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
			"fromDate":  "01/04/2023",
			"toDate":    "16/04/2024",
		},
	}
	bodyReq, _ := json.Marshal(timoBody)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyReq))
	token := "eyJraWQiOiJjNmIzMTNlNS1iNTAzLTQzMzEtYTdiMi0zYmNiNDVjOTA2YjUiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIzb0hEZDhwTGlEdkdrS2pFR2FweG93IiwidWlkIjoiYzc1ZThjMDItM2E0MS00Y2JlLTgzMjAtZTBiZjYzZDQ5YjBiIn0.c7c9wEhNBhcSU4524GR6cdHF9XfFFN-SWIoLosS93Do"
	if err != nil {
		log.Printf("Error when create new request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("token", token)
	bankResponse := &ResponseData{}
	resRaw, err := client.Do(request)

	if err != nil {
		log.Printf("Error when send requst. %v", err)
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
