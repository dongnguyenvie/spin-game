package ginuser

import (
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	userbiz "nolan/spin-game/modules/users/biz"
	userrepo "nolan/spin-game/modules/users/repository"
	userstorage "nolan/spin-game/modules/users/storage"

	"github.com/gin-gonic/gin"
)

func GetNonce(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		walletAddrs := c.Param("addrs")

		db := ctx.GetMaiDBConnection()
		storage := userstorage.NewSQLStore(db)
		userRepo := userrepo.NewUserNonceRepo(storage)
		userNonceBiz := userbiz.NewGetUserNonceBiz(userRepo)
		nonce := userNonceBiz.GetUserNonce(c.Request.Context(), walletAddrs)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]interface{}{"nonce": nonce}))
	}
}
