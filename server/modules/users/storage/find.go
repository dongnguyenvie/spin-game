package userstorage

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	var data usermodel.User
	db := s.db.Table(usermodel.User{}.TableName())

	if err := db.Where(condition).First(&data); err != nil {
		return nil, common.RecordNotFound
	}

	return &data, nil
}
