package categorystore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/category/categorymodel"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(categorymodel.Category{}.TableName()).Where("id = ?", id).Updates(
		map[string]interface{}{
			"status": 0,
		},
	).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
