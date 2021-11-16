package pbtransaction

import (
	"context"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	transactionbiz "nolan/spin-game/modules/transactions/biz"
	transactionmodel "nolan/spin-game/modules/transactions/model"
	transactionstorage "nolan/spin-game/modules/transactions/storage"
)

func DepositSuccessfully(appctx appctx.AppContext) {
	pb := appctx.GetPubsub()

	ch, _ := pb.Subscribe(context.Background(), common.ChannelTxDepositSuccess)

	db := appctx.GetMaiDBConnection()
	txStorage := transactionstorage.NewSQLStore(db)
	// txRepo := transactionrepo.NewTransactionRepo(txStorage)
	txBiz := transactionbiz.NewDepositingBiz(txStorage)

	go func() {
		defer common.AppRecover()
		for {
			txId := (<-ch).Data().(int)
			err := txBiz.FinishDepositingTx(context.Background(), txId, transactionmodel.Successfully)
			if err != nil {
				panic(err)
			}
		}
	}()
}
