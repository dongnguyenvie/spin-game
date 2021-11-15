package walletrepo

import (
	"context"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

type WalletStorage interface {
	ListDataWithCondition(ctx context.Context, filter *walletmodel.Filter, paging *common.Paging, moreKeys ...string) ([]walletmodel.Wallet, error)
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*walletmodel.Wallet, error)
	Create(ctx context.Context, data *walletmodel.WalletCreate) error
	UpdateBalance(ctx context.Context, condition map[string]interface{}, amount int) error
}

type walletRepo struct {
	storage WalletStorage
}

func NewWalletRepo(storage WalletStorage) *walletRepo {
	return &walletRepo{storage: storage}
}
