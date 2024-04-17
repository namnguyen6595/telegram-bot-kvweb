package bank

type BankInterface interface {
	GetTransaction() ([]*TransactionResponse, error)
}

func InitialBanks(bankType string) BankInterface {
	switch bankType {
	case "timo":
		return &TimoBank{}
	default:
		return nil
	}
}
