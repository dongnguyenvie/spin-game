package walletbiz

import (
	"context"
	"errors"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

type WithdrawStorage interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*walletmodel.Wallet, error)
	Update_balance(ctx context.Context, condition map[string]interface{}, amount int) error
}

type withdrawBiz struct {
	walletStorage WithdrawStorage
}

func NewWithdraw(walletStorage WithdrawStorage) *withdrawBiz {
	return &withdrawBiz{walletStorage: walletStorage}
}

func (biz *withdrawBiz) Withdraw(ctx context.Context, user *tokenprovider.TokenPayload, data *walletmodel.Withdraw) error {
	wallet, _ := biz.walletStorage.FindDataWithCondition(ctx, map[string]interface{}{"user_id": user.UserId})

	if wallet == nil {
		return common.ErrEntityExisted(walletmodel.EntityName, errors.New("wallet is not found"))
	}

	if data.Amount <= 0 {
		return common.ErrInsufficientFunds(errors.New("amount must be greather than 0"))
	}

	if wallet.Balance <= data.Amount {
		return common.ErrInsufficientFunds(errors.New("wallet is out of money"))
	}

	err := biz.walletStorage.Update_balance(ctx, map[string]interface{}{"id": wallet.Id}, -data.Amount)

	if err != nil {
		return err
	}

	return nil
}
