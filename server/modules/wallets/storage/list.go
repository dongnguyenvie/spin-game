package walletstorage

import (
	"context"
	"nolan/spin-game/components/common"
	walletmodel "nolan/spin-game/modules/wallets/model"
)

func (s *sqlStore) ListDataWithCondition(ctx context.Context, filter *walletmodel.Filter, paging *common.Paging, moreKeys ...string) ([]walletmodel.Wallet, error) {
	var result []walletmodel.Wallet

	db := s.db.Table(walletmodel.Wallet{}.TableName())

	err := db.Find(&result).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
