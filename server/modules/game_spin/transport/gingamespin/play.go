package gingamespin

import (
	"fmt"
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	"nolan/spin-game/components/tokenprovider"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"

	"github.com/gin-gonic/gin"
)

func Play(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data gamespinmodel.GameSpinPlay
		user := c.MustGet(common.CurrentUser).(*tokenprovider.TokenPayload)
		fmt.Printf("%v", user)

		data.RandomAward()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
