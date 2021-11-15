package walletbiz

import (
	"context"
	"errors"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

type DepositWalletRepo interface {
	FindWallet(ctx context.Context, conditions map[string]interface{}) (*walletmodel.Wallet, error)
	UpdateBalance(ctx context.Context, condition map[string]interface{}, amount int) error
}

type depositWalletBiz struct {
	walletRepo DepositWalletRepo
}

func NewDepositBiz(walletRepo DepositWalletRepo) *depositWalletBiz {
	return &depositWalletBiz{walletRepo: walletRepo}
}

func (biz *depositWalletBiz) DepositWallet(ctx context.Context, walletId int, amount int) error {
	wallet, _ := biz.walletRepo.FindWallet(ctx, map[string]interface{}{"id": walletId})

	if wallet != nil {
		return common.ErrEntityExisted(walletmodel.EntityName, errors.New("wallet is exists"))
	}

	err := biz.walletRepo.UpdateBalance(ctx, map[string]interface{}{"id": walletId}, amount)

	if err != nil {
		return err
	}

	return nil
}
