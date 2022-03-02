package categorymodel

import (
	"golang-blog-api/common"
	"strings"

	"github.com/gosimple/slug"
)

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Slug            string `json:"slug" gorm:"column:slug;"`
	Description     string `json:"description" gorm:"column:description;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (res *CategoryCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.Slug = strings.TrimSpace(res.Slug)

	if len(res.Name) == 0 {
		return ErrCategoryNameCannotBeEmpty
	}

	if len(res.Slug) == 0 {
		return ErrCategorySlugCannotBeEmpty
	}

	if !slug.IsSlug(res.Slug) {
		return ErrCategorySlugIsInvalid
	}

	return nil
}
