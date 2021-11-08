package userrepo

import (
	"context"
	"fmt"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

func (userRepo *userRepo) UpdateUser(ctx context.Context, id string, data map[string]interface{}) error {
	fmt.Println("id", id)
	fmt.Println("data", data)
	if err := userRepo.storage.Update(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
