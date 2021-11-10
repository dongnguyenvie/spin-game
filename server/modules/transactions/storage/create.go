package transactionstorage

import (
	"context"
	"nolan/spin-game/components/common"
	transactionmodel "nolan/spin-game/modules/transactions/model"
)

func (s *sqlStore) Create(ctx context.Context, data *transactionmodel.Transaction) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
