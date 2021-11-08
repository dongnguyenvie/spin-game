package walletstorage

import (
	"context"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

func (s *sqlStore) Update(ctx context.Context, condition map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(walletmodel.Wallet{}.TableName())

	db = db.Where(condition)

	if err := db.Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
