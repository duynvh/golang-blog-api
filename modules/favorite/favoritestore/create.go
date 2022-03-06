package favoritestore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *favoritemodel.FavoriteCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
