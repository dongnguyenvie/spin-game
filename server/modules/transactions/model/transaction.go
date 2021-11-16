package transactionmodel

import (
	"nolan/spin-game/components/common"
)

const EntityName = "Transactions"

type Transaction struct {
	common.SQLModel `json:",inline"`
	Tx              string `json:"tx" gorm:"column:tx;"`
	Type            int    `json:"type" gorm:"column:type;"`
	Credit          int64  `json:"credit" gorm:"column:credit;"`
	Debit           int    `json:"debit" gorm:"column:debit;"`
	CreateBy        *int   `json:"createBy,omitempty" gorm:"column:create_by"`
	Package         int    `json:"package" gorm:"column:package"`
	Status          int    `json:"status" gorm:"column:status;default:1;"`

	UserId   int `json:"userId" gorm:"column:user_id;"`
	WalletId int `json:"walletId" gorm:"column:wallet_id;"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (u *Transaction) GetId() int {
	return u.Id
}

// var (
// 	Deposit    = 0
// 	BuyPackage = 1
// )

var (
	Waiting      = 3
	Successfully = 4
	Failed       = 5
)
