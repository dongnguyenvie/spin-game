package gamespinbiz

import (
	"context"
	"nolan/spin-game/components/tokenprovider"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
)

type MyPackageStorage interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*gamespinmodel.GameSpin, error)
	Update_package(ctx context.Context, condition map[string]interface{}, amount int) error
}

type myPackageBiz struct {
	gamespinStorage MyPackageStorage
}

func NewMyPackageBiz(gamespinStorage GameSpinStorage) *myPackageBiz {
	return &myPackageBiz{gamespinStorage: gamespinStorage}
}

func (biz *myPackageBiz) MyPackage(ctx context.Context, user *tokenprovider.TokenPayload) (*gamespinmodel.GameSpin, error) {
	gamespinmodel, err := biz.gamespinStorage.FindDataWithCondition(ctx, map[string]interface{}{"user_id": user.UserId})

	if err != nil {
		return nil, err
	}

	return gamespinmodel, nil
}
