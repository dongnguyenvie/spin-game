package pbtransaction

import (
	"context"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/pubsub"

	transactionbiz "nolan/spin-game/modules/transactions/biz"
	transactionmodel "nolan/spin-game/modules/transactions/model"
	transactionstorage "nolan/spin-game/modules/transactions/storage"
)

func RequestDeposit(appctx appctx.AppContext) {
	pb := appctx.GetPubsub()

	ch, _ := pb.Subscribe(context.Background(), common.ChannelDepositBC)

	db := appctx.GetMaiDBConnection()
	txstorage := transactionstorage.NewSQLStore(db)
	// txRepo := transactionrepo.NewTransactionRepo(txstorage)
	txBiz := transactionbiz.NewDepositingBiz(txstorage)

	go func() {
		defer common.AppRecover()

		for {
			userDeposit := (<-ch).Data().(common.UserDeposit)
			newTx := transactionmodel.Transaction{
				Tx:       userDeposit.Tx,
				Credit:   userDeposit.Amount.Int64(),
				Debit:    0,
				WalletId: userDeposit.WalletId,
				Type:     common.DepositType,
				Status:   transactionmodel.Waiting,
			}
			tx, err := txBiz.StartDepositingTx(context.Background(), &newTx)

			if err != nil {
				panic("create tx failed")
			}

			newTxEvent := pubsub.NewMessage(common.NewTx{
				Id:       tx.Id,
				Type:     tx.Type,
				Status:   tx.Status,
				Debit:    0,
				Credit:   int(tx.Credit),
				WalletId: tx.WalletId,
			})
			pb.Publish(context.Background(), common.ChannelNewTx, newTxEvent)

		}
	}()

}
