package transactionstorage

import (
	"context"
	"nolan/spin-game/components/common"
	transactionmodel "nolan/spin-game/modules/transactions/model"
)

func (s *sqlStore) ListDataWithCondition(ctx context.Context, filter *transactionmodel.Filter, paging *common.Paging, moreKeys ...string) ([]transactionmodel.Transaction, error) {
	var result []transactionmodel.Transaction

	db := s.db.Table(transactionmodel.Transaction{}.TableName())

	err := db.Find(&result).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
