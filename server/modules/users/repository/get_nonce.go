package userrepo

import (
	"context"
	"nolan/spin-game/components/common"
)

// type UserNonceStorage interface {
// 	FindDataWithCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
// }

// type userNonceRepo struct {
// 	storage UserNonceStorage
// }

// func NewUserNonceRepo(storage UserNonceStorage) *userNonceRepo {
// 	return &userNonceRepo{
// 		storage: storage,
// 	}
// }

func (repo *userRepo) GetUserNonce(ctx context.Context, address string) (int, error) {
	user, err := repo.storage.FindDataWithCondition(ctx, map[string]interface{}{
		"wallet_address": address,
	})

	if err != nil {
		return -1, common.ErrCannotGetEntity("user nonce", err)
	}

	return user.Nonce, nil
}
