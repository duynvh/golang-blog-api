package component

import (
	"golang-blog-api/component/uploadprovider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secret string
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secret string) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secret: secret}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secret
}
