package gamespinstorage

import (
	"context"
	"nolan/spin-game/components/common"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
	usermodel "nolan/spin-game/modules/users/model"
)

func (s *sqlStore) ListDataWithCondition(ctx context.Context, filter *usermodel.Filter, paging *common.Paging, moreKeys ...string) ([]gamespinmodel.GameSpin, error) {
	var result []gamespinmodel.GameSpin

	db := s.db.Table(gamespinmodel.GameSpin{}.TableName())

	err := db.Find(&result).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
