package gamespinbiz

import (
	"context"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
)

type CreateGameSpinStorage interface {
	Create(ctx context.Context, data *gamespinmodel.GameSpin) error
}

type createGameSpinBiz struct {
	storage CreateGameSpinStorage
}

func NewCreateGameSpinBiz(storage CreateGameSpinStorage) *createGameSpinBiz {
	return &createGameSpinBiz{storage: storage}
}

func (biz *createGameSpinBiz) CreateGameSpin(ctx context.Context, data *gamespinmodel.GameSpin) (*gamespinmodel.GameSpin, error) {
	err := biz.storage.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
