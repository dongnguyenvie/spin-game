package transactionstorage

import (
	"context"
	"nolan/spin-game/components/common"
	transactionmodel "nolan/spin-game/modules/transactions/model"
)

func (s *sqlStore) Update(ctx context.Context, condition map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(transactionmodel.Transaction{}.TableName())

	db = db.Where(condition)

	if err := db.Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
