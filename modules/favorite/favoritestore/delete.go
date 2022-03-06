package favoritestore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
)

func (s *sqlStore) Delete(ctx context.Context, userId, postId int) error {
	db := s.db

	if err := db.Table(favoritemodel.Favorite{}.TableName()).
		Where("user_id = ?", userId).
		Where("post_id = ?", postId).Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
