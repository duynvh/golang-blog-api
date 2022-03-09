package userbiz

import (
	"context"
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/component/tokenprovider"
	"golang-blog-api/modules/user/usermodel"

	"go.opencensus.io/trace"
)

type LoginStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type loginBiz struct {
	appCtx        component.AppContext
	store         LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(store LoginStore, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{store: store, tokenProvider: tokenProvider, hasher: hasher, expiry: expiry}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)
func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	ctx1, span1 := trace.StartSpan(ctx, "user.biz.login")
	user, err := biz.store.FindUser(ctx1, map[string]interface{}{"email": data.Email})
	span1.End()

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	_, span2 := trace.StartSpan(ctx, "user.biz.login.gen-jwt")
	passHashed := biz.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		span2.End()
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Roke:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	span2.End()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
