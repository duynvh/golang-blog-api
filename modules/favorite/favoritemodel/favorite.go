package favoritemodel

import (
	"golang-blog-api/common"
	"golang-blog-api/modules/post/postmodel"
	"time"
)

const EntityName = "Favorite"

type Favorite struct {
	CreatedAt *time.Time         `json:"created_at" gorm:"created_at"`
	UserId    int                `json:"-" gorm:"column:user_id;"`
	PostId    int                `json:"-" gorm:"column:post_id;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
	Post      *postmodel.Post    `json:"post" gorm:"preload:false;"`
}

func (Favorite) TableName() string {
	return "favorites"
}

// func (data *Category) Mask(isAdmin bool) {
// 	data.GenUID(common.DBTypeCategory)
// }

var (
	ErrFavoritePostIsMissing = common.NewCustomError(nil, "missing post", "ErrFavoritePostIsMissing")
	ErrFavoriteAPostTwice    = common.NewCustomError(nil, "you favorited this post", "ErrFavoriteAPostTwice")
)

// use error function when you want to capture the root error
func ErrFavoritePostIsInvalid(err error) *common.AppError {
	return common.NewCustomError(err, "invalid post", "ErrFavoritePostIsInvalid")
}
