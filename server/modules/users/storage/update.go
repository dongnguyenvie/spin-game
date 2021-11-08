package userstorage

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

func (s *sqlStore) Update(ctx context.Context, condition map[string]interface{}, data map[string]interface{}) error {
	db := s.db.Table(usermodel.User{}.TableName())

	db = db.Where(condition)

	if err := db.Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
