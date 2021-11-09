package walletmodel

import (
	"nolan/spin-game/components/common"
)

const EntityName = "Transactions"

type Transaction struct {
	common.SQLModel `json:",inline"`
	Tx              string `json:"tx" gorm:"column:balance;"`
	Type            int    `json:"type" gorm:"column:balance;"`
	Credit          int    `json:"credit" gorm:"column:balance;"`
	Debit           int    `json:"debit" gorm:"column:status;"`
	CreateBy        *int   `json:"create_by" gorm:"column:user_id"`
	Status          int    `json:"status" gorm:"column:user_id"`
	Package         int    `json:"package" gorm:"column:user_id"`

	UserId   int `json:"user_id" gorm:"column:balance;"`
	WalletId int `json:"wallet_id" gorm:"column:balance;"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (u *Transaction) GetId() int {
	return u.Id
}

var ()
