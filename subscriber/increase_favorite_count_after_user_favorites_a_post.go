package subscriber

import (
	"context"
	"golang-blog-api/component"
	"golang-blog-api/modules/post/poststore"
	"golang-blog-api/pubsub"
)

type HasPostId interface {
	GetPostId() int
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
