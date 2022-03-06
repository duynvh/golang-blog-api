package favoritebiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
	"golang-blog-api/modules/post/postmodel"
)

type FavoriteStore interface {
	Create(ctx context.Context, data *favoritemodel.FavoriteCreate) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*favoritemodel.Favorite, error)
}

type PostStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type favoriteBiz struct {
	store     FavoriteStore
	postStore PostStore
}

func NewFavoriteBiz(store FavoriteStore, postStore PostStore) *favoriteBiz {
	return &favoriteBiz{store: store, postStore: postStore}
}

func (biz *favoriteBiz) Favorite(
	ctx context.Context,
	data *favoritemodel.FavoriteCreate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// favorite
	{
		_, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"user_id": data.UserId, "post_id": data.PostId})
		if err != common.ErrRecordNotFound {
			return favoritemodel.ErrFavoriteAPostTwice
		}
	}

	// post
	{
		post, err := biz.postStore.FindDataByCondition(ctx, map[string]interface{}{"id": data.PostId})
		if err != nil || post.Status == 0 {
			return favoritemodel.ErrFavoritePostIsInvalid(err)
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(favoritemodel.EntityName, err)
	}

	return nil
}
