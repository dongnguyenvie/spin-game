package userrepo

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

type CreateUserStorage interface {
	FindDataWithCondition(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserRepo struct {
	storage CreateUserStorage
}

func NewCreateUserRepo(storage CreateUserStorage) *createUserRepo {
	return &createUserRepo{storage: storage}
}

func (userRepo *createUserRepo) FindUser(ctx context.Context, condition map[string]interface{}) (*usermodel.User, error) {
	result, err := userRepo.storage.FindDataWithCondition(ctx, condition)
	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}
	return result, nil
}

func (userRepo *createUserRepo) CreateUser(ctx context.Context, data *usermodel.UserCreate) (*usermodel.UserCreate, error) {
	if err := userRepo.storage.Create(ctx, data); err != nil {
		return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return data, nil
}
