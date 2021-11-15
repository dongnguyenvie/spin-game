package walletstorage

import (
	"context"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateBalance(ctx context.Context, condition map[string]interface{}, amount int) error {
	db := s.db.Table(walletmodel.Wallet{}.TableName())

	db = db.Where(condition)

	if err := db.Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
