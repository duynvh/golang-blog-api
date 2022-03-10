package postbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
	"log"
)

type ListStore interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		filter *postmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type UserFavoritedStore interface {
	GetFavoriteCountOfPosts(ctx context.Context, ids []int) (map[int]int, error)
}

type listBiz struct {
	store              ListStore
	userFavoritedStore UserFavoritedStore
}

func NewListBiz(store ListStore, userFavoritedStore UserFavoritedStore) *listBiz {
	return &listBiz{store: store, userFavoritedStore: userFavoritedStore}
}

func (biz *listBiz) List(
	ctx context.Context,
	filter *postmodel.Filter,
	paging *common.Paging,
) ([]postmodel.Post, error) {
	result, err := biz.store.List(ctx, nil, filter, paging, "User", "Category")

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	ids := make([]int, len(result))

	for i := range result {
		ids[i] = result[i].Id
	}

	mapResFavorited, err := biz.userFavoritedStore.GetFavoriteCountOfPosts(ctx, ids)

	if err != nil {
		log.Println("Cannot get post favorited:", err)
	}

	for i, item := range result {
		result[i].FavoriteCount = mapResFavorited[item.Id]
	}

	return result, nil
}
