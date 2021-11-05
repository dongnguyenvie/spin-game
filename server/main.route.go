package main

import (
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/modules/users/transport/ginuser"
	"nolan/spin-game/modules/wallets/transport/ginwallet"

	"github.com/gin-gonic/gin"
)

func setupMainRoute(appCtx appctx.AppContext, group *gin.RouterGroup) {

	{
		group.POST("/signin", ginuser.SignIn(appCtx))
		group.POST("/signup", ginuser.SignUp(appCtx))
		group.GET("/profile", ginuser.Profile(appCtx))
	}

	{
		userGr := group.Group("/users")
		userGr.GET("", ginuser.ListUser(appCtx))
	}

	{
		walletGr := group.Group("/wallets")
		walletGr.GET("", ginwallet.ListWallet(appCtx))
		walletGr.POST("", ginwallet.CreateWallet(appCtx))
		walletGr.GET("/:id", ginwallet.GetWallet(appCtx))
		walletGr.DELETE("/:id", ginwallet.DeleteWallet(appCtx))
	}

}
