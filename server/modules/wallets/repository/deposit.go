package walletrepo

import (
	"context"
)

func (repo *walletRepo) UpdateBalance(ctx context.Context, condition map[string]interface{}, amount int) error {
	err := repo.storage.UpdateBalance(ctx, condition, amount)
	if err != nil {
		return err
	}
	return nil
}
