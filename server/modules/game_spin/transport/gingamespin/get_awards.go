package gingamespin

import (
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"
	gamespinmodel "nolan/spin-game/modules/game_spin/model"

	"github.com/gin-gonic/gin"
)

func GetAwards(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		awards := gamespinmodel.GameSpinPlay{}.Awards()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]interface{}{
			"awards": awards,
		}))
	}
}
