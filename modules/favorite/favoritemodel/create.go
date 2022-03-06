package favoritemodel

import (
	"time"
)

type FavoriteCreate struct {
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UserId    int        `json:"-" gorm:"column:user_id;"`
	PostId    int        `json:"-" gorm:"column:post_id;"`
}

func (FavoriteCreate) TableName() string {
	return Favorite{}.TableName()
}

func (f *FavoriteCreate) Validate() error {
	if f.PostId == 0 {
		return ErrFavoritePostIsMissing
	}

	return nil
}

func (f *FavoriteCreate) GetPostId() int {
	return f.PostId
}

func (f *FavoriteCreate) GetUserId() int {
	return f.UserId
}
