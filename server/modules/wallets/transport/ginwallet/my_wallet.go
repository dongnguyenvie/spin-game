package ginwallet

import (
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	walletbiz "nolan/spin-game/modules/wallets/biz"
	walletstorage "nolan/spin-game/modules/wallets/storage"

	"github.com/gin-gonic/gin"
)

func MyWallet(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(*tokenprovider.TokenPayload)

		db := appctx.GetMaiDBConnection()
		walletStorage := walletstorage.NewSQLStore(db)
		walletBiz := walletbiz.NewMyWalletBiz(walletStorage)
		result, err := walletBiz.MyWallet(c, map[string]interface{}{"user_id": user.UserId})

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
