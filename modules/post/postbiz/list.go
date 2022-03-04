package postbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
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

type listBiz struct {
	store ListStore
}

func NewListBiz(store ListStore) *listBiz {
	return &listBiz{store: store}
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

	return result, nil
}
