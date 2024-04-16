package bank

type TransactionResponse struct {
	Items    *[]TransactionItem `json:"item"`
	DispDate string             `json:"dispDate"`
}

type TransactionItem struct {
	TxnType                 string        `json:"txnType"`
	RawTxnType              string        `json:"rawTxnType"`
	TxnDesc                 string        `json:"txnDesc"`
	TxnTime                 string        `json:"txnTime"`
	TxnTimeTimestamp        string        `json:"txnTimeTimestamp"`
	TxnAmount               int           `json:"txnAmount"`
	TxnAmountMC             float64       `json:"txnAmountMC"`
	RemainingAmount         float64       `json:"remainingAmount"`
	TxnNarrative            string        `json:"txnNarrative"`
	PayeeName               string        `json:"payeeName"`
	RefNo                   string        `json:"refNo"`
	GoalSavePlanName        string        `json:"goalSavePlanName"`
	CreditAccount           string        `json:"creditAccount"`
	BankXID                 string        `json:"bankXID"`
	TimoAppType             string        `json:"timoAppType"`
	TimoDesc1               string        `json:"timoDesc1"`
	TimoDesc2               string        `json:"timoDesc2"`
	Cur                     string        `json:"cur"`
	BulkId                  int           `json:"bulkId"`
	BulkRows                int           `json:"bulkRows"`
	BulkStatus              int           `json:"bulkStatus"`
	TransactionTime         string        `json:"transactionTime"`
	TimoMigration           bool          `json:"timoMigration"`
	CanChangeReactionEmoji  bool          `json:"canChangeReactionEmoji"`
	OriginalFullDescription string        `json:"originalFullDescription"`
	IsInternalTransfer      int           `json:"isInternalTransfer"`
	IsReactionEnabled       bool          `json:"isReactionEnabled"`
	TxnAccountDisplayName   string        `json:"txnAccountDisplayName"`
	TxnBankName             string        `json:"txnBankName"`
	TxnBankAccount          string        `json:"txnBankAccount"`
	IsMoneyTransferTxn      bool          `json:"isMoneyTransferTxn"`
	TxnProductType          string        `json:"txnProductType"`
	ChannelCode             string        `json:"channelCode"`
	PfmLabels               []interface{} `json:"pfmLabels"`
	PfmLabelsDisplayable    bool          `json:"pfmLabelsDisplayable"`
	PfmLabelsEditable       bool          `json:"pfmLabelsEditable"`
	CategoryCode            string        `json:"categoryCode"`
	IsTimoMem               bool          `json:"isTimoMem"`
	TxnTitle                string        `json:"txnTitle"`
}
