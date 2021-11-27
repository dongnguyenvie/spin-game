package gingamespin

import (
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	gamespinbiz "nolan/spin-game/modules/game_spin/biz"
	gamespinstorage "nolan/spin-game/modules/game_spin/storage"
	walletstorage "nolan/spin-game/modules/wallets/storage"

	"github.com/gin-gonic/gin"
)

func Play(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(*tokenprovider.TokenPayload)
		db := appctx.GetMaiDBConnection()
		gamespinStorage := gamespinstorage.NewSQLStore(db)
		walletStorage := walletstorage.NewSQLStore(db)
		biz := gamespinbiz.NewPlaySpinGame(gamespinStorage, walletStorage)

		result, err := biz.Play(c, user)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
