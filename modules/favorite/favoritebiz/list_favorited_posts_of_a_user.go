package favoritebiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
	"golang-blog-api/modules/post/postmodel"
)

type listFavoritedPostOfAuthUserBiz struct {
	store ListStore
}

func NewListFavoritedPostOfAuthUserBiz(store ListStore) *listFavoritedPostOfAuthUserBiz {
	return &listFavoritedPostOfAuthUserBiz{store: store}
}

func (biz *listFavoritedPostOfAuthUserBiz) ListFavoritedPostsOfAUser(
	ctx context.Context,
	filter *favoritemodel.Filter,
	paging *common.Paging,
) ([]*postmodel.Post, error) {
	result, err := biz.store.List(ctx, nil, filter, paging, "Post")
	if err != nil {
		return nil, common.ErrCannotListEntity(favoritemodel.EntityName, err)
	}

	posts := make([]*postmodel.Post, len(result))
	for i, item := range result {
		if item.Post == nil {
			continue
		}
		posts[i] = item.Post
		posts[i].CreatedAt = item.CreatedAt
		posts[i].UpdatedAt = nil
	}

	return posts, nil
}
