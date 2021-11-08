package userbiz

import (
	"context"
	"errors"
	"fmt"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
	"strconv"
)

type SigninRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
	UpdateUser(ctx context.Context, id string, data map[string]interface{}) error
}

type signinBiz struct {
	signinRepo SigninRepo
}

func NewSigninBiz(signinRepo SigninRepo) *signinBiz {
	return &signinBiz{signinRepo: signinRepo}
}

func (biz *signinBiz) Signin(ctx context.Context, data *usermodel.GetUserByWallet) (*usermodel.User, error) {
	user, err := biz.signinRepo.FindUser(ctx, map[string]interface{}{"wallet_address": data.WalletAddress})

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, common.ErrEntityNotFound(usermodel.User{}.TableName(), errors.New("wallet address is not found"))
	}

	data.MessageSigning = strconv.Itoa(user.Nonce)

	if err := data.IsSignatureValid(); err != nil {
		return nil, err
	}

	go func() {
		defer common.AppRecover()
		err := biz.signinRepo.UpdateUser(ctx, strconv.Itoa(user.Id), map[string]interface{}{
			"nonce": common.RandomNonce(),
		})
		fmt.Println("err", err)
	}()

	return user, nil
}
