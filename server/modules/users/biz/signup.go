package userbiz

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

type SignupRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type signupBiz struct {
	signupRepo SignupRepo
}

func NewRegisterBusiness(signupRepo SignupRepo) *signupBiz {
	return &signupBiz{
		signupRepo: signupRepo,
	}
}

func (biz *signupBiz) Signup(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.signupRepo.FindUser(ctx, map[string]interface{}{"walletAddress": data.WalletAddress})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	if err := biz.signupRepo.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil

}
