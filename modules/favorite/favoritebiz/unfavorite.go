package favoritebiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
)

type UnfavoriteStore interface {
	Delete(ctx context.Context, userId, postId int) error
}

type unfavoriteBiz struct {
	store UnfavoriteStore
}

func NewUnfavoriteBiz(store UnfavoriteStore) *unfavoriteBiz {
	return &unfavoriteBiz{store: store}
}

func (biz *unfavoriteBiz) Unfavorite(
	ctx context.Context,
	userId int,
	postId int,
) error {

	if err := biz.store.Delete(ctx, userId, postId); err != nil {
		return common.ErrCannotDeleteEntity(favoritemodel.EntityName, err)
	}

	return nil
}
