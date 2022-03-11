package categorybiz

import (
	"context"
	"errors"
	"golang-blog-api/common"
	"golang-blog-api/modules/category/categorymodel"
	"testing"
)

type mockCreateStore struct{}

func (s *mockCreateStore) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if data.Name == "abc def" {
		return errors.New("abc def is invalid in db")
	}
	return nil
}

type testingItem struct {
	Input  categorymodel.CategoryCreate
	Actual error
	Expect error
}

func TestCreate(t *testing.T) {
	biz := NewCreateBiz(&mockCreateStore{})

	dataTable := []testingItem{
		{
			Input: categorymodel.CategoryCreate{
				Name: "",
			},
			Expect: common.ErrInvalidRequest(categorymodel.ErrCategoryNameCannotBeEmpty),
		},
		{
			Input: categorymodel.CategoryCreate{
				Name: "abc",
				Slug: "",
			},
			Expect: common.ErrInvalidRequest(categorymodel.ErrCategorySlugCannotBeEmpty),
		},
		{
			Input: categorymodel.CategoryCreate{
				Name: "abc",
				Slug: "abc123 ádds",
			},
			Expect: common.ErrInvalidRequest(categorymodel.ErrCategorySlugIsInvalid),
		},
		{
			Input: categorymodel.CategoryCreate{
				Name: "abc def",
				Slug: "abc-def",
			},
			Expect: common.ErrCannotCreateEntity(categorymodel.EntityName, errors.New("abc def is invalid in db")),
		},
		{
			Input: categorymodel.CategoryCreate{
				Name: "abc def 123",
				Slug: "abc-def-123",
			},
			Expect: nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.Create(context.Background(), &item.Input)

		if actual == nil {
			if item.Expect != nil {
				t.Errorf("Expect error is %s but actual is %v\n", item.Expect, actual)
			}
			continue
		}

		if actual.Error() != item.Expect.Error() {
			t.Errorf("Expect error is %s but actual is %v\n", item.Expect, actual)
		}
	}
}
