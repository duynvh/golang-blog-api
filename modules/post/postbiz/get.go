package postbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
)

type GetStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type getBiz struct {
	store GetStore
}

func NewGetBiz(store GetStore) *getBiz {
	return &getBiz{store: store}
}

func (biz *getBiz) Get(
	ctx context.Context,
	id int,
) (*postmodel.Post, error) {
	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id}, "User", "Category")

	if err != nil {
		if err != common.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	return result, nil
}
