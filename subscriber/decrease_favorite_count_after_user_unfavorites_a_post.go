package subscriber

import (
	"context"
	"golang-blog-api/component"
	"golang-blog-api/modules/post/poststore"
	"golang-blog-api/pubsub"

	"go.opencensus.io/trace"
)

func RunDecreaseUnfavoriteCountAfterUserFavoritesAPost(appCtx component.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease favorite count after user unfavorites a post",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			// _ = message.Data().([]int)[0] // simulate crashes
			store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
			postId := message.Data().(int)

			ctx1, span := trace.StartSpan(ctx, "pubsub.sub.RunDecreaseUnfavoriteCountAfterUserFavoritesAPost")
			defer span.End()
			return store.DecreaseFavoriteCount(ctx1, postId)
		},
	}
}
