package usermodel

import "nolan/spin-game/components/common"

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
