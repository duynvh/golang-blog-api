package postbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/category/categorymodel"
	"golang-blog-api/modules/post/postmodel"
)

type UpdateStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	Update(ctx context.Context, id int, data *postmodel.PostUpdate) error
}

type updateBiz struct {
	store UpdateStore
	categoryStore CategoryStore
}

func NewUpdateBiz(store UpdateStore, categoryStore CategoryStore) *updateBiz {
	return &updateBiz{store: store, categoryStore: categoryStore}
}

func (biz *updateBiz) Update(ctx context.Context, id int, data *postmodel.PostUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if data.Slug != "" {
		_, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"slug": data.Slug}) 
		if data.Slug != oldData.Slug && err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(postmodel.EntityName, nil)
		}
	}

	{
		category, err := biz.categoryStore.FindDataByCondition(ctx, map[string]interface{}{"id": data.CategoryId})
		if category == nil {
			return common.ErrEntityNotFound(categorymodel.EntityName, err)
		}

		if category.Status == 0 {
			return common.ErrEntityDeleted(categorymodel.EntityName, nil)
		}
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}

	return nil
}
