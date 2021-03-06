package favoritebiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
)

type ListStore interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		filter *favoritemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]favoritemodel.Favorite, error)
}

type listUsersFavoritedAPostBiz struct {
	store ListStore
}

func NewListUsersFavoritedAPostBiz(store ListStore) *listUsersFavoritedAPostBiz {
	return &listUsersFavoritedAPostBiz{store: store}
}

func (biz *listUsersFavoritedAPostBiz) ListUserFavoritedAPost(
	ctx context.Context,
	filter *favoritemodel.Filter,
	paging *common.Paging,
) ([]*common.SimpleUser, error) {
	result, err := biz.store.List(ctx, nil, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(favoritemodel.EntityName, err)
	}

	users := make([]*common.SimpleUser, len(result))
	for i, item := range result {
		if item.User == nil {
			continue
		}

		users[i] = item.User
		users[i].CreatedAt = item.CreatedAt
		users[i].UpdatedAt = nil
	}

	return users, nil
}
