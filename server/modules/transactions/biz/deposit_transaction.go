package transactionbiz

import (
	"context"
	transactionmodel "nolan/spin-game/modules/transactions/model"
)

type DepositRepo interface{}

type depositBiz struct {
	depositRepo DepositRepo
}

func NewDepositBiz(depositRepo DepositRepo) *depositBiz {
	return &depositBiz{depositRepo: depositRepo}
}

func (biz *depositBiz) DepositTransaction(ctx context.Context, data *transactionmodel.TransactionDeposit) (*transactionmodel.Transaction, error) {
	// TODO: create tsx with tx from blockchain
	return nil, nil
}

func (biz *depositBiz) UpdateTransactionStatus(ctx context.Context, data *transactionmodel.Transaction) error {
	// TODO: update transaction[failed][successfully] when the process is finished
	return nil
}
