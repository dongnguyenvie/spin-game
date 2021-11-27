package walletmodel

import (
	"nolan/spin-game/components/common"
)

const EntityName = "Wallet"

type Wallet struct {
	common.SQLModel `json:",inline"`
	Balance         int `json:"balance" gorm:"column:balance;"`
	Status          int `json:"status" gorm:"column:status;"`
	UserId          int `json:"userId" gorm:"column:user_id"`
}

func (Wallet) TableName() string {
	return "wallets"
}

func (u *Wallet) GetId() int {
	return u.Id
}

func (u *Wallet) Transfer(toAddress string) int {
	return u.Id
}

var ()
