package poststore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *postmodel.PostCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
