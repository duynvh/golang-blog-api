package component

import (
	"golang-blog-api/component/uploadprovider"
	"golang-blog-api/pubsub"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
	GetGRPCClientConnection() grpc.ClientConnInterface
	SetGRPCClientConnection(grpc.ClientConnInterface)
}

type appCtx struct {
	db                   *gorm.DB
	upProvider           uploadprovider.UploadProvider
	secret               string
	pb                   pubsub.Pubsub
	grpcClientConnection grpc.ClientConnInterface
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

func (ctx *appCtx) GetGRPCClientConnection() grpc.ClientConnInterface {
	return ctx.grpcClientConnection
}

func (ctx *appCtx) SetGRPCClientConnection(c grpc.ClientConnInterface) {
	ctx.grpcClientConnection = c
}
