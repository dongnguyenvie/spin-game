package gamespinstorage

import (
	"context"
	"nolan/spin-game/components/common"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Update_package(ctx context.Context, condition map[string]interface{}, amount int) error {
	db := s.db.Table(gamespinmodel.GameSpin{}.TableName())

	db = db.Where(condition)

	if err := db.Update("package", gorm.Expr("package + ?", amount)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
