package pbwallet

import (
	"context"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	walletbiz "nolan/spin-game/modules/wallets/biz"
	walletrepo "nolan/spin-game/modules/wallets/repository"
	walletstorage "nolan/spin-game/modules/wallets/storage"
)

func DepositWallet(appctx appctx.AppContext) {
	pb := appctx.GetPubsub()

	ch, _ := pb.Subscribe(context.Background(), common.ChannelNewTx)

	db := appctx.GetMaiDBConnection()
	walletStorage := walletstorage.NewSQLStore(db)
	walletRepo := walletrepo.NewWalletRepo(walletStorage)
	depositBiz := walletbiz.NewDepositBiz(walletRepo)

	go func() {
		defer common.AppRecover()
		for {
			tx := (<-ch).Data().(common.NewTx)
			if tx.Type == common.DepositType {
				amount := tx.Credit + tx.Debit
				err := depositBiz.DepositWallet(context.Background(), tx.WalletId, amount)
				if err != nil {
					panic(err)
				}
				// TODO: create publish message update transaction status after wallet is updated
			}
		}
	}()
}
