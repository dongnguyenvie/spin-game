package usermodel

import (
	"errors"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/ethsigner"
)

type GetUserByWallet struct {
	WalletAddress  string `json:"walletAddress"`
	MessageSigning string `json:"-" gorm:"-"`
	Signature      string `json:"signature" gorm:"-"`
}

func (u *GetUserByWallet) IsSignatureValid() error {
	signers := ethsigner.NewSigner()
	isVerifying := signers.VerifySig(u.WalletAddress, u.Signature, []byte(u.MessageSigning))

	if !isVerifying {
		return common.ErrInvalidRequest(errors.New("signature is invalid"))
	}

	return nil
}
