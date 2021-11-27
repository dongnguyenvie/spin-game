package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const (
	CurrentUser = "user"
)

const (
	ChannelCreateUser       = "ChannelCreateUser"
	ChannelDepositBC        = "ChannelDepositBC"
	ChannelNewTx            = "ChannelNewTx"
	ChannelTxDepositSuccess = "ChannelDepositSuccess"
)

var (
	DepositType    = 0
	BuyPackageType = 1
)

const (
	StatusLocked = 0
	StatusActive = 1
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}

const (
	PackagePrice    = 100000000000
	PackageQuantity = 5
)

const (
	CoinDecimal = 1e18
)
