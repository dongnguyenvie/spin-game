package userstorage

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

func (s *sqlStore) ListDataWithCondition(ctx context.Context, filter *usermodel.Filter, paging *common.Paging, moreKeys ...string) ([]usermodel.User, error) {
	var result []usermodel.User

	db := s.db.Table(usermodel.User{}.TableName())

	err := db.Find(&result).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
