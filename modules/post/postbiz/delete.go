package postbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
)

type DeleteStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	Delete(ctx context.Context, id int) error
}

type deleteBiz struct {
	store DeleteStore
}

func NewDeleteBiz(store DeleteStore) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) Delete(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(postmodel.EntityName, err)
	}

	return nil
}
