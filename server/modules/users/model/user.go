package usermodel

import (
	"errors"
	"nolan/spin-game/components/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	WalletAddress   string `json:"walletAddress" gorm:"column:wallet_address"`
	Password        string `json:"-" gorm:"column:password;"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
