package subscriber

import (
	"context"
	"golang-blog-api/component"
	"golang-blog-api/modules/post/poststore"
	"golang-blog-api/pubsub"
	"golang-blog-api/skio"
)

type HasPostId interface {
	GetPostId() int
	GetUserId() int
}

func RunIncreaseFavoriteCountAfterUserFavoritesAPost(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase favorite count after user favorites a post",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
			favoriteData := message.Data().(HasPostId)
			return store.IncreaseFavoriteCount(ctx, favoriteData.GetPostId())
		},
	}
}

func EmitRealtimeAfterUserFavoritesAPost(appCtx component.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Emit realtime after user favorite post",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasPostId)
			return rtEngine.EmitToUser(likeData.GetUserId(), string(message.Channel()), likeData)
		},
	}
}