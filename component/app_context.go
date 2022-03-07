package component

import (
	"golang-blog-api/component/uploadprovider"
	"golang-blog-api/pubsub"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secret     string
	pb         pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secret string, pb pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secret: secret, pb: pb}
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

func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
	return ctx.pb
}
