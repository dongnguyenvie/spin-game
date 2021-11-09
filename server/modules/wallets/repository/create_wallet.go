package walletrepo

import (
	"context"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

func (repo *walletRepo) FindWallet(ctx context.Context, condition map[string]interface{}) (*walletmodel.Wallet, error) {
	result, err := repo.storage.FindDataWithCondition(ctx, condition)
	if err != nil {
		return nil, common.ErrCannotListEntity(walletmodel.EntityName, err)
	}
	return result, nil
}

func (repo *walletRepo) CreateWallet(ctx context.Context, data *walletmodel.WalletCreate) (*walletmodel.WalletCreate, error) {
	if err := repo.storage.Create(ctx, data); err != nil {
		return nil, common.ErrCannotCreateEntity(walletmodel.EntityName, err)
	}
	return data, nil
}
