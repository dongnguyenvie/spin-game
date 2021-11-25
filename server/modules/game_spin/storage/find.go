package gamespinstorage

import (
	"context"
	"nolan/spin-game/components/common"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*gamespinmodel.GameSpin, error) {
	var data gamespinmodel.GameSpin
	db := s.db.Table(gamespinmodel.GameSpin{}.TableName())

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, common.RecordNotFound
	}

	return &data, nil
}
