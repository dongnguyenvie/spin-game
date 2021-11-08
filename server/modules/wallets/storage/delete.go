package walletstorage

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	db := s.db.Table(usermodel.User{}.TableName())

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
