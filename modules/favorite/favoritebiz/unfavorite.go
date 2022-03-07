package favoritebiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
	"golang-blog-api/pubsub"
)

type UnfavoriteStore interface {
	Delete(ctx context.Context, userId, postId int) error
}

type unfavoriteBiz struct {
	store  UnfavoriteStore
	pubsub pubsub.Pubsub
}

func NewUnfavoriteBiz(store UnfavoriteStore, pubsub pubsub.Pubsub) *unfavoriteBiz {
	return &unfavoriteBiz{store: store, pubsub: pubsub}
}

func (biz *unfavoriteBiz) Unfavorite(
	ctx context.Context,
	userId int,
	postId int,
) error {

	if err := biz.store.Delete(ctx, userId, postId); err != nil {
		return common.ErrCannotDeleteEntity(favoritemodel.EntityName, err)
	}

	// side effect
	// go func() {
	// 	defer common.AppRecover()
	// 	job := asyncjob.NewJob(func(ctx context.Context) error {
	// 		return biz.decStore.DecreaseFavoriteCount(ctx, postId)
	// 	})

	// 	_ = asyncjob.NewGroup(true, job).Run(ctx)
	// }()
	biz.pubsub.Publish(ctx, common.TopicUserUnfavoritePost, pubsub.NewMessage(postId))

	return nil
}
