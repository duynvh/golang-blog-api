package ginuser

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/component/hasher"
	"golang-blog-api/component/tokenprovider/jwt"
	"golang-blog-api/modules/user/userbiz"
	"golang-blog-api/modules/user/usermodel"
	"golang-blog-api/modules/user/userstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		store := userstore.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)

		data, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
