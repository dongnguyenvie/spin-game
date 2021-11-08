package appctx

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	GetSecretKey() string
	GetEthClient() *ethclient.Client
	GetSmartContractAddr() string
}

type appCtx struct {
	db        *gorm.DB
	secret    string
	ethClient *ethclient.Client
	smAddrs   string
}

func NewAppContext(db *gorm.DB, ethClient *ethclient.Client, secretKey string, smartContractAddrs string) *appCtx {
	return &appCtx{
		db:        db,
		ethClient: ethClient,
		secret:    secretKey,
		smAddrs:   smartContractAddrs,
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
