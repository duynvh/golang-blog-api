package favoritestore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/favorite/favoritemodel"
)

type SqlData struct {
	PostId        int `gorm:column:"post_id";`
	FavoriteCount int `gorm:"column:"favorite_count;"`
}

func (s *sqlStore) GetFavoriteCountOfPosts(
	ctx context.Context,
	postIds []int,
) (map[int]int, error) {
	result := make(map[int]int)
	var listFavorite []SqlData
	db := s.db

	if err := db.Table(favoritemodel.Favorite{}.TableName()).
		Select("post_id, count(post_id) as favorite_count").
		Where("post_id in (?)", postIds).
		Group("post_id").
		Find(&listFavorite).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listFavorite {
		result[item.PostId] = item.FavoriteCount
	}

	return result, nil
}
