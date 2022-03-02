package categorymodel

import (
	"golang-blog-api/common"
	"strings"

	"github.com/gosimple/slug"
)

type CategoryUpdate struct {
	common.SQLUpdateModel `json:",inline"`
	Name                  string `json:"name" gorm:"column:name;"`
	Slug                  string `json:"slug" gorm:"column:slug;"`
	Description           string `json:"description" gorm:"column:description;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

func (res *CategoryUpdate) Validate() error {
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
