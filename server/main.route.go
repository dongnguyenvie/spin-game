package main

import (
	"nolan/spin-game/components/appctx"

	"github.com/gin-gonic/gin"
)

func setupMainRoute(appCtx appctx.AppContext, group *gin.RouterGroup) {
	group.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "1",
		})
	})
}
