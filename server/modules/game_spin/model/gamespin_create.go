package gamespinmodel

import (
	"nolan/spin-game/components/common"
)

type GameSpinCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int `json:"userId" gorm:"column:user_id"`
	Package         int `json:"package" gorm:"column:package;"`
	Status          int `json:"status" gorm:"column:status;"`
}
