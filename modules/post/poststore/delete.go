package poststore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).Updates(
		map[string]interface{}{
			"status": 0,
		},
	).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
