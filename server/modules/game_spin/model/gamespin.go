package gamespinmodel

import (
	"nolan/spin-game/components/common"
)

const EntityName = "GameSpin"

type GameSpin struct {
	common.SQLModel `json:",inline"`
	UserId          int `json:"userId" gorm:"column:user_id"`
	Package         int `json:"package" gorm:"column:package;"`
	Status          int `json:"status" gorm:"column:status;"`
}

func (GameSpin) TableName() string {
	return "game_spins"
}

func (u *GameSpin) GetId() int {
	return u.Id
}
