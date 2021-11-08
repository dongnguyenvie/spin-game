package userrepo

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

type UserStorage interface {
	ListDataWithCondition(ctx context.Context, filter *usermodel.Filter, paging *common.Paging, moreKeys ...string) ([]usermodel.User, error)
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
	Delete(context context.Context, id int) error
	Update(ctx context.Context, condition map[string]interface{}, data map[string]interface{}) error
}

type userRepo struct {
	storage UserStorage
}

func NewUserRepo(storage UserStorage) *userRepo {
	return &userRepo{storage: storage}
}
