package pbwallet

import (
	"context"
	"fmt"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	usermodel "nolan/spin-game/modules/users/model"
	walletbiz "nolan/spin-game/modules/wallets/biz"
	walletmodel "nolan/spin-game/modules/wallets/model"
	walletrepo "nolan/spin-game/modules/wallets/repository"
	walletstorage "nolan/spin-game/modules/wallets/storage"
)

func CreateWallet(appctx appctx.AppContext) {

	pb := appctx.GetPubsub()
	userCreateSub, _ := pb.Subscribe(context.Background(), common.ChannelCreateUser)
	db := appctx.GetMaiDBConnection()
	walletStorage := walletstorage.NewSQLStore(db)
	walletRepo := walletrepo.NewWalletRepo(walletStorage)
	walletBiz := walletbiz.NewSigninBiz(walletRepo)

	go func() {
		defer common.AppRecover()

		for {
			userCreated := (<-userCreateSub).Data().(*usermodel.UserCreate)
			walletCreate := walletmodel.WalletCreate{
				Balance: 0,
				UserId:  userCreated.Id,
				Status:  common.StatusActive,
			}
			wallet, err := walletBiz.CreateWallet(context.Background(), &walletCreate)

			if err != nil {
				panic(err)
			}

			fmt.Printf("create wallet:: userid:%d, walletId:%d \n", userCreated.Id, wallet.Id)
		}
	}()
}
