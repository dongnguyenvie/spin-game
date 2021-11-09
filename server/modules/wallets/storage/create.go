package walletstorage

import (
	"context"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

func (s *sqlStore) Create(ctx context.Context, data *walletmodel.WalletCreate) error {
	db := s.db.Begin()

	if err := db.Table(walletmodel.Wallet{}.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
