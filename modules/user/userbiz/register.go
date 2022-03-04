package userbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/modules/user/usermodel"
)

type RegisterStore interface {
	Create(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStore
	hasher Hasher
}

func NewRegisterBiz(store RegisterStore, hasher Hasher) *registerBiz {
	return &registerBiz{store: store, hasher: hasher}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
