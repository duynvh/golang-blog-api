package middleware

import (
	"context"
	"errors"
	"fmt"
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/component/tokenprovider/jwt"
	"golang-blog-api/modules/user/usermodel"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	// "Authorzation:" "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we user user_id to find from BD
func RequireAuth(appCtx component.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		// db := appCtx.GetMainDBConnection()
		// store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		ctx, span := trace.StartSpan(c.Request.Context(), "middleware.RequiredAuth")
		user, err := authStore.FindUser(ctx, map[string]interface{}{"id": payload.UserId})
		span.End()

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
