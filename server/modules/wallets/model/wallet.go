package walletmodel

import (
	"nolan/spin-game/components/common"
)

const EntityName = "Wallet"

type Wallet struct {
	common.SQLModel `json:",inline"`
	Balance         string `json:"balance" gorm:"column:balance;"`
	Status          int    `json:"status" gorm:"column:status;"`
}

func (Wallet) TableName() string {
	return "wallets"
}

func (u *Wallet) GetId() int {
	return u.Id
}

var ()
