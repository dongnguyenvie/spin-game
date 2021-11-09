package walletstorage

import (
	"context"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*walletmodel.Wallet, error) {
	var data walletmodel.Wallet
	db := s.db.Table(usermodel.User{}.TableName())

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, common.RecordNotFound
	}

	return &data, nil
}
