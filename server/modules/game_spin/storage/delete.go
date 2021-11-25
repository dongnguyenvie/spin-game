package gamespinstorage

import (
	"context"
	"nolan/spin-game/components/common"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	db := s.db.Table(gamespinmodel.GameSpin{}.TableName())

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
