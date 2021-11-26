package pbgamespin

import (
	"context"
	"fmt"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	gamespinbiz "nolan/spin-game/modules/game_spin/biz"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"
	gamespinstorage "nolan/spin-game/modules/game_spin/storage"
	usermodel "nolan/spin-game/modules/users/model"
)

func CreateGameSpin(appctx appctx.AppContext) {

	pb := appctx.GetPubsub()
	userCreateSub, _ := pb.Subscribe(context.Background(), common.ChannelCreateUser)
	db := appctx.GetMaiDBConnection()
	gamespinStorage := gamespinstorage.NewSQLStore(db)
	gamespinBiz := gamespinbiz.NewCreateGameSpinBiz(gamespinStorage)

	go func() {
		defer common.AppRecover()

		for {
			userCreated := (<-userCreateSub).Data().(*usermodel.UserCreate)
			gamespinCreate := gamespinmodel.GameSpin{
				UserId: userCreated.Id,
				Status: common.StatusActive,
			}
			gamespin, err := gamespinBiz.CreateGameSpin(context.Background(), &gamespinCreate)

			if err != nil {
				panic(err)
			}

			fmt.Printf("create gamespin:: userid:%d, gamespin:%d \n", userCreated.Id, gamespin.Id)
		}
	}()
}
