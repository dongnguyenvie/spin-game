package usermodel

import (
	"nolan/spin-game/components/common"

	"github.com/ethereum/go-ethereum/crypto"
)

type UserCreate struct {
	common.SQLModel `json:",inline"`
	User            `json:",inline"`
	//
	Password       string `json:"password" gorm:"column:password"`
	MessageSigning string `json:"messageSign" gorm:"-"`
	Signature      string `json:"signature" gorm:"-"`
}

func (u *UserCreate) IsSignatureValid() bool {
	data := []byte(u.MessageSigning)
	hash := crypto.Keccak256Hash(data)

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), []byte(u.Signature))
	if err != nil || string(sigPublicKey[:]) != u.WalletAddress {
		return false
	}

	return true
}
