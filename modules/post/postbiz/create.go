package postbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/category/categorymodel"
	"golang-blog-api/modules/post/postmodel"
)

type CreateStore interface {
	Create(ctx context.Context, data *postmodel.PostCreate) error
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
}

type CategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type createBiz struct {
	store         CreateStore
	categoryStore CategoryStore
}

func NewCreateBiz(store CreateStore, categoryStore CategoryStore) *createBiz {
	return &createBiz{store: store, categoryStore: categoryStore}
}

func (biz *createBiz) Create(ctx context.Context, data *postmodel.PostCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	{
		_, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"slug": data.Slug})
		if err != common.ErrRecordNotFound {
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

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	return nil
}
