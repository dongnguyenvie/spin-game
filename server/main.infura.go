package main

import (
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/modules/wallets/transport/bcwallet"
)

func setupInfura(appCtx appctx.AppContext) {
	bcwallet.Deposit(appCtx)
	// bctransaction.Deposit(appCtx)
}
