package appctx

import (
	"nolan/spin-game/components/pubsub"

	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	GetSecretKey() string
	GetEthClient() *ethclient.Client
	GetSmartContractAddr() string
	GetPubsub() pubsub.PubSub
}

type appCtx struct {
	db        *gorm.DB
	secret    string
	ethClient *ethclient.Client
	smAddrs   string
	pubsub    pubsub.PubSub
}

func NewAppContext(db *gorm.DB, ethClient *ethclient.Client, pubsub pubsub.PubSub, secretKey string, smartContractAddrs string) *appCtx {
	return &appCtx{
		db:        db,
		ethClient: ethClient,
		secret:    secretKey,
		smAddrs:   smartContractAddrs,
		pubsub:    pubsub,
	}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetEthClient() *ethclient.Client {
	return ctx.ethClient
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secret
}

func (ctx *appCtx) GetSmartContractAddr() string {
	return ctx.smAddrs
}

func (ctx *appCtx) GetPubsub() pubsub.PubSub {
	return ctx.pubsub
}
