package walletbiz

import (
	"context"
	"errors"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

type CreateWalletRepo interface {
	FindWallet(ctx context.Context, conditions map[string]interface{}) (*walletmodel.Wallet, error)
	CreateWallet(ctx context.Context, data *walletmodel.WalletCreate) (*walletmodel.WalletCreate, error)
}

type createWalletBiz struct {
	walletRepo CreateWalletRepo
}

func NewCreateWalletBiz(walletRepo CreateWalletRepo) *createWalletBiz {
	return &createWalletBiz{walletRepo: walletRepo}
}

func (biz *createWalletBiz) CreateWallet(ctx context.Context, data *walletmodel.WalletCreate) (*walletmodel.WalletCreate, error) {
	wallet, _ := biz.walletRepo.FindWallet(ctx, map[string]interface{}{"user_id": data.UserId})

	if wallet != nil {
		return nil, common.ErrEntityExisted(walletmodel.EntityName, errors.New("wallet is exists"))
	}

	result, err := biz.walletRepo.CreateWallet(ctx, data)

	if err != nil {
		return nil, err
	}

	return result, nil
}
