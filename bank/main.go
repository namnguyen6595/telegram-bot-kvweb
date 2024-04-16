package bank

type BankInterface interface {
	GetTransaction() ([]*TransactionResponse, error)
}

func InitialBanks() BankInterface {
	return &TimoBank{}
}
