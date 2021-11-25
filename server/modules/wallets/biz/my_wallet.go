package walletbiz

import (
	"context"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

type MyWalletStorage interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*walletmodel.Wallet, error)
}

type myWalletBiz struct {
	storage MyWalletStorage
}

func NewMyWalletBiz(storage MyWalletStorage) *myWalletBiz {
	return &myWalletBiz{storage: storage}
}

func (biz *myWalletBiz) MyWallet(ctx context.Context, data map[string]interface{}) (*walletmodel.Wallet, error) {
	result, err := biz.storage.FindDataWithCondition(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
