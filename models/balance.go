package models

import (
	"gorm.io/gorm"
	"kvweb-bot/constant"
)

type BalanceDmsClient struct {
	Db *gorm.DB
}

func (h *BalanceDmsClient) GetAllTransaction(req *constant.TransactionsRequest) {
	//result := h.db.Get("")
}
