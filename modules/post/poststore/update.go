package poststore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *postmodel.PostUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseFavoriteCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).
				Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
					return common.ErrDB(err)
				}
	return nil
}

func (s *sqlStore) DecreaseFavoriteCount(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).
				Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
					return common.ErrDB(err)
				}
	return nil
}