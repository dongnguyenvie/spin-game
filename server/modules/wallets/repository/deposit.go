package walletrepo

import (
	"context"
)

func (repo *walletRepo) Deposit(ctx context.Context, condition map[string]interface{}, amount int) error {
	err := repo.storage.Deposit(ctx, condition, amount)
	if err != nil {
		return err
	}
	return nil
}
