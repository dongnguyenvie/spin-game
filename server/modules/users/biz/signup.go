package userbiz

import (
	"context"
	usermodel "nolan/spin-game/modules/users/model"
)

type Hasher interface {
	Hash(data string) string
}

type SignupRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserCreate, error)
}

type signupBiz struct {
	signupRepo SignupRepo
	hasher     Hasher
}

func NewSignupBiz(signupRepo SignupRepo, hasher Hasher) *signupBiz {
	return &signupBiz{
		signupRepo: signupRepo,
		hasher:     hasher,
	}
}

func (biz *signupBiz) Signup(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserCreate, error) {
	user, err := biz.signupRepo.FindUser(ctx, map[string]interface{}{"wallet_address": data.WalletAddress})

	if user != nil {
		return nil, err
	}

	data.Password = biz.hasher.Hash(data.Password)
	result, err := biz.signupRepo.CreateUser(ctx, data)
	if err != nil {
		return nil, err
	}

	return result, nil

}
