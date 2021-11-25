package gamespinstorage

import (
	"context"
	"nolan/spin-game/components/common"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
)

func (s *sqlStore) Update(ctx context.Context, condition map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(gamespinmodel.GameSpin{}.TableName())

	db = db.Where(condition)

	if err := db.Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
