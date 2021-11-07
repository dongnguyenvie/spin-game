package usermodel

import (
	"errors"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/ethsigner"
)

type UserCreate struct {
	User `json:",inline"`
	//
	Password       string `json:"password" gorm:"column:password"`
	MessageSigning string `json:"messageSign" gorm:"-"`
	Signature      string `json:"signature" gorm:"-"`
}

func (u *UserCreate) IsSignatureValid() error {
	signers := ethsigner.NewSigner()
	isVerifying := signers.VerifySig(u.WalletAddress, u.Signature, []byte(u.MessageSigning))

	if !isVerifying {
		return common.ErrInvalidRequest(errors.New("signature is invalid"))
	}

	return nil
}
