package transactionbiz

import (
	"context"
	"errors"
	"nolan/spin-game/components/common"
	transactionmodel "nolan/spin-game/modules/transactions/model"
)

type DepositRepo interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*transactionmodel.Transaction, error)
	Create(ctx context.Context, data *transactionmodel.Transaction) error
	Update(ctx context.Context, condition map[string]interface{}, data map[string]interface{}) error
}

type depositBiz struct {
	depositRepo DepositRepo
}

func NewDepositingBiz(depositRepo DepositRepo) *depositBiz {
	return &depositBiz{depositRepo: depositRepo}
}

func (biz *depositBiz) StartDepositingTx(ctx context.Context, data *transactionmodel.Transaction) (*transactionmodel.Transaction, error) {
	tx, _ := biz.depositRepo.FindDataWithCondition(ctx, map[string]interface{}{"tx": data.Tx})

	if tx != nil {
		return nil, common.ErrEntityExisted("transaction", errors.New("tx is exists"))
	}

	err := biz.depositRepo.Create(ctx, data)

	if err != nil {
		return data, nil
	}

	return data, nil
}

func (biz *depositBiz) FinishDepositingTx(ctx context.Context, id int, status int) error {

	err := biz.depositRepo.Update(ctx, map[string]interface{}{"id": id}, map[string]interface{}{"status": status})
	if err != nil {
		return err
	}

	return nil
}
