package main

import (
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/middleware"
	"nolan/spin-game/modules/game_spin/transport/gingamespin"
	"nolan/spin-game/modules/transactions/transport/gintransaction"
	"nolan/spin-game/modules/users/transport/ginuser"
	"nolan/spin-game/modules/wallets/transport/ginwallet"

	"github.com/gin-gonic/gin"
)

func setupMainRoute(appCtx appctx.AppContext, group *gin.RouterGroup) {

	{
		group.POST("/signin", ginuser.SignIn(appCtx))                            // done
		group.POST("/signup", ginuser.SignUp(appCtx))                            // done
		group.GET("/profile", middleware.Guard(appCtx), ginuser.Profile(appCtx)) // done
	}

	{
		userGr := group.Group("/users")
		userGr.GET("", ginuser.ListUser(appCtx))
		userGr.GET("/nonce/:addrs", ginuser.GetNonce(appCtx))
	}

	{
		walletGr := group.Group("/wallets")
		walletGr.GET("/me", middleware.Guard(appCtx), ginwallet.MyWallet(appCtx)) // done
		walletGr.GET("", ginwallet.ListWallet(appCtx))
		walletGr.POST("", ginwallet.CreateWallet(appCtx))
		walletGr.GET("/:id", ginwallet.GetWallet(appCtx))
		walletGr.DELETE("/:id", ginwallet.DeleteWallet(appCtx))
	}

	{
		transactionGr := group.Group("/transactions")
		transactionGr.GET("", gintransaction.ListTransaction(appCtx))
		transactionGr.GET("/:id", gintransaction.GetTransaction(appCtx))
	}

	{
		gameGr := group.Group("/spin")
		gameGr.GET("/histories", middleware.Guard(appCtx), gingamespin.ListHistory(appCtx))
		gameGr.POST("/play", middleware.Guard(appCtx), gingamespin.Play(appCtx))
		gameGr.POST("/buy-package", middleware.Guard(appCtx), gingamespin.BuyPackage(appCtx))

	}

}
