package userbiz

import (
	"context"
)

type UserNonceRepo interface {
	GetUserNonce(ctx context.Context, address string) (int, error)
}

type getUserNonceBiz struct {
	getUserNonceRepo UserNonceRepo
}

func NewGetUserNonceBiz(getUserNonceRepo UserNonceRepo) *getUserNonceBiz {
	return &getUserNonceBiz{getUserNonceRepo: getUserNonceRepo}
}

func (biz *getUserNonceBiz) GetUserNonce(ctx context.Context, address string) int {

	nonce, err := biz.getUserNonceRepo.GetUserNonce(ctx, address)

	if err != nil {
		panic(err)
	}

	return nonce

}
