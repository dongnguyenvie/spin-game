package main

import (
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/modules/users/transport/bcuser"
)

func setupInfura(appCtx appctx.AppContext) {
	bcuser.Deposit(appCtx)
}
