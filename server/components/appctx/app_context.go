package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	GetSecretKey() string
}

type appCtx struct {
	db     *gorm.DB
	secret string
}

func NewAppContext(db *gorm.DB, secret string) *appCtx {
	return &appCtx{
		db:     db,
		secret: secret,
	}
}

func (ctx *appCtx) GetMaiDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.secret
}
