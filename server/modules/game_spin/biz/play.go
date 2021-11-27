package gamespinbiz

import (
	"context"
	"errors"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
)

type GameSpinStorage interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*gamespinmodel.GameSpin, error)
	Update_package(ctx context.Context, condition map[string]interface{}, amount int) error
}

type GameSpinWalletStorage interface {
	Update_balance(ctx context.Context, condition map[string]interface{}, amount int) error
}

type gameSpinBiz struct {
	gamespinStorage GameSpinStorage
	walletStorage   GameSpinWalletStorage
}

func NewPlaySpinGame(storage GameSpinStorage, walletStorage GameSpinWalletStorage) *gameSpinBiz {
	return &gameSpinBiz{gamespinStorage: storage, walletStorage: walletStorage}
}

func (biz *gameSpinBiz) Play(ctx context.Context, user *tokenprovider.TokenPayload) (*gamespinmodel.GameSpinPlay, error) {
	var data gamespinmodel.GameSpinPlay

	gamespin, err := biz.gamespinStorage.FindDataWithCondition(ctx, map[string]interface{}{"user_id": user.UserId})
	if err != nil {
		return nil, err
	}
	if gamespin.Package < 1 {
		return nil, common.ErrInsufficientFunds(errors.New("insufficient package"))
	}
	err = biz.gamespinStorage.Update_package(ctx, map[string]interface{}{"user_id": user.UserId}, -1)
	if err != nil {
		return nil, err
	}
	data.RandomAward()
	award := common.CoinDecimal * data.Award

	err = biz.walletStorage.Update_balance(ctx, map[string]interface{}{"user_id": user.UserId}, int(award))

	if err != nil {
		return nil, err
	}

	return &data, nil
}
