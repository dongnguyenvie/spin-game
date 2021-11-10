package transactionstorage

import (
	"context"
	"nolan/spin-game/components/common"
	transactionmodel "nolan/spin-game/modules/transactions/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	db := s.db.Table(transactionmodel.Transaction{}.TableName())

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
