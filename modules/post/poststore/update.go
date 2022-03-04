package poststore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *postmodel.PostUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
