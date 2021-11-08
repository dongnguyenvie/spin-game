package ginuser

import (
	"fmt"
	"net/http"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/common"

	"github.com/gin-gonic/gin"
)

func Profile(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser)
		fmt.Println("user request", user)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
