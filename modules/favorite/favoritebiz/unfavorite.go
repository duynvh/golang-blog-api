package favoritebiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/component/asyncjob"
	"golang-blog-api/modules/favorite/favoritemodel"
)

type UnfavoriteStore interface {
	Delete(ctx context.Context, userId, postId int) error
}

type DecreaseFavoriteCountStore interface {
	DecreaseFavoriteCount(ctx context.Context, id int) error
}

type unfavoriteBiz struct {
	store    UnfavoriteStore
	decStore DecreaseFavoriteCountStore
}

func NewUnfavoriteBiz(store UnfavoriteStore, decStore DecreaseFavoriteCountStore) *unfavoriteBiz {
	return &unfavoriteBiz{store: store, decStore: decStore}
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
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decStore.DecreaseFavoriteCount(ctx, postId)
		})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
