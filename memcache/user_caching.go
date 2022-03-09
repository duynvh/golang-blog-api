package memcache

import (
	"context"
	"fmt"
	"golang-blog-api/modules/user/usermodel"
)

type RealStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type userCaching struct {
	store Caching
	realStore RealStore
}

func NewUserCaching(store Caching, realStore RealStore) *userCaching {
	return &userCaching{
		store: store,
		realStore: realStore,
	}
}

func (uc *userCaching) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*usermodel.User, error) {
	userId := conditions["id"].(int)
	key := fmt.Sprintf("user-%d", userId)

	userInCache := uc.store.Read(key)

	if userInCache != nil {
		return userInCache.(*usermodel.User), nil
	}

	user, err := uc.realStore.FindUser(ctx, conditions, moreKeys...)
	if err != nil {
		return nil, err
	}

	uc.store.Write(key, user)
	return user, nil
}
