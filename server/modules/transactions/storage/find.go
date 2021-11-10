package transactionstorage

import (
	"context"
	"nolan/spin-game/components/common"
	transactionmodel "nolan/spin-game/modules/transactions/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*transactionmodel.Transaction, error) {
	var data transactionmodel.Transaction
	db := s.db.Table(transactionmodel.Transaction{}.TableName())

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, common.RecordNotFound
	}

	return &data, nil
}
