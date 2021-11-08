package userrepo

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

// type ListUserStorage interface {
// 	ListDataWithCondition(context context.Context, filter *usermodel.Filter, paging *common.Paging, moreKeys ...string) ([]usermodel.User, error)
// }

// type listUserRepo struct {
// 	storage ListUserStorage
// }

// func NewListUserRepo(storage ListUserStorage) *listUserRepo {
// 	return &listUserRepo{storage: storage}
// }

func (userRepo *userRepo) ListUser(context context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error) {
	result, err := userRepo.storage.ListDataWithCondition(context, filter, paging, "")
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return result, nil
}
