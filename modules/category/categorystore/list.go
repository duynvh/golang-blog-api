package categorystore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/category/categorymodel"
)

func (s *sqlStore) List(
	ctx context.Context,
	conditions map[string]interface{},
	filter *categorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]categorymodel.Category, error) {
	var result []categorymodel.Category
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(categorymodel.Category{}.TableName()).
		Where(conditions).
		Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
