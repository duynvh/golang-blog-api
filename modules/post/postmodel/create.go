package postmodel

import (
	"golang-blog-api/common"
	"strings"

	"github.com/gosimple/slug"
)

type PostCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Title           string        `json:"title" gorm:"column:title;"`
	Slug            string        `json:"slug" gorm:"column:slug;"`
	ShortDesc       string        `json:"short_desc" gorm:"column:short_desc;"`
	Body            string        `json:"body" gorm:"column:body;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
	Keywords        string        `json:"keywords" gorm:"column:keywords;"`
	CategoryId      int           `json:"-" gorm:"column:category_id;"`
	FakeCategoryId  *common.UID   `json:"category_id" gorm:"-"`
	UserId          int           `json:"-" gorm:"column:owner_id;"`
}

func (PostCreate) TableName() string {
	return Post{}.TableName()
}

func (p *PostCreate) Fulfill() {
	if p.FakeCategoryId != nil {
		p.CategoryId = int(p.FakeCategoryId.GetLocalID())
	}
}

func (res *PostCreate) Validate() error {
	res.Title = strings.TrimSpace(res.Title)
	res.Slug = strings.TrimSpace(res.Slug)
	res.ShortDesc = strings.TrimSpace(res.ShortDesc)
	res.Body = strings.TrimSpace(res.Body)
	res.Keywords = strings.TrimSpace(res.Keywords)

	if len(res.Title) == 0 {
		return ErrPostTitleCannotBeEmpty
	}

	if len(res.Slug) == 0 {
		return ErrPostSlugCannotBeEmpty
	}

	if !slug.IsSlug(res.Slug) {
		return ErrPostSlugIsInvalid
	}

	if len(res.Body) == 0 {
		return ErrPostBodyCannotBeEmpty
	}

	if res.FakeCategoryId == nil {
		return ErrPostCategoryCannotBeEmpty
	}

	if len(res.Keywords) > 255 {
		return ErrPostKeywordsIsTooLong
	}

	return nil
}

func (data *PostCreate) Mask(isAdmin bool) {
	data.GenUID(common.DBTypePost)
}
