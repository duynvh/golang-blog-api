package userstore

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/user/usermodel"

	"go.opencensus.io/trace"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	_, span := trace.StartSpan(ctx, "store.user.find-user")
	defer span.End()
	
	var user usermodel.User
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
