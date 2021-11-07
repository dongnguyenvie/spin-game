package userbiz

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

type ListUserRepo interface {
	ListUser(context context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error)
}

type listUserBiz struct {
	repo ListUserRepo
}

func NewListUserBiz(repo ListUserRepo) *listUserBiz {
	return &listUserBiz{repo: repo}
}

func (biz *listUserBiz) ListUser(ctx context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error) {
	result, err := biz.repo.ListUser(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
