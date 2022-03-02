package categorybiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/category/categorymodel"
)

type GetStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
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
) (*categorymodel.Category, error) {
	result, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	return result, nil
}
