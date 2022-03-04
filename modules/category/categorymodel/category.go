package categorymodel

import "golang-blog-api/common"

const EntityName = "Category"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Slug            string `json:"slug" gorm:"column:slug;"`
	Description     string `json:"description" gorm:"column:description;"`
}

func (Category) TableName() string {
	return "categories"
}

func (data *Category) Mask(isAdmin bool) {
	data.GenUID(common.DBTypeCategory)
}

var (
	ErrCategoryNameCannotBeEmpty = common.NewCustomError(nil, "category name can't be blank", "ErrCategoryNameCannotBeEmpty")
	ErrCategoryNameIsTooLong     = common.NewCustomError(nil, "category name is too long", "ErrCategoryNameIsTooLong")
	ErrCategorySlugCannotBeEmpty = common.NewCustomError(nil, "slug can't be blank", "ErrCategoryNameCannotBeEmpty")
	ErrCategorySlugIsTooLong     = common.NewCustomError(nil, "slug is too long", "ErrCategoryNameIsTooLong")
	ErrCategorySlugIsInvalid     = common.NewCustomError(nil, "slug is invalid", "ErrCategorySlugIsInvalid")
)
