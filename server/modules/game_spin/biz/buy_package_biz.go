package gamespinbiz

import (
	"context"
	"errors"
	"fmt"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

type WalletStorage interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*walletmodel.Wallet, error)
	Update_balance(ctx context.Context, condition map[string]interface{}, amount int) error
}
type BuyPackageStorage interface {
	Update_package(ctx context.Context, condition map[string]interface{}, amount int) error
}

type buyPackageBiz struct {
	walletStorage     WalletStorage
	buyPackageStorage BuyPackageStorage
}

func NewBuyPackageBiz(walletStorage WalletStorage, buyPackageStorage BuyPackageStorage) *buyPackageBiz {
	return &buyPackageBiz{walletStorage: walletStorage, buyPackageStorage: buyPackageStorage}
}

func (biz *buyPackageBiz) BuyPackage(ctx context.Context, user *tokenprovider.TokenPayload, data *gamespinmodel.BuyPackage) (*gamespinmodel.GameSpin, error) {
	wallet, _ := biz.walletStorage.FindDataWithCondition(ctx, map[string]interface{}{"user_id": user.UserId})

	fmt.Println("user.UserId", user.UserId)

	if wallet == nil {
		return nil, common.ErrEntityExisted(walletmodel.EntityName, errors.New("wallet is not exists"))
	}

	if wallet.Balance < common.PackagePrice {
		return nil, common.ErrInsufficientPackage(errors.New("insufficient funds"))
	}

	err := biz.walletStorage.Update_balance(ctx, map[string]interface{}{"user_id": user.UserId}, -(common.PackagePrice * data.Quantity))

	if err != nil {
		return nil, err
	}

	err = biz.buyPackageStorage.Update_package(ctx, map[string]interface{}{"user_id": user.UserId}, data.Quantity)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
