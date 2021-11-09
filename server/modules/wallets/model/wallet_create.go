package walletmodel

import "nolan/spin-game/components/common"

type WalletCreate struct {
	common.SQLModel `json:",inline"`
	Balance         int `json:"balance" gorm:"column:balance;"`
	Status          int `json:"status" gorm:"column:status;"`
	UserId          int `json:"userId" gorm:"column:user_id"`
}
