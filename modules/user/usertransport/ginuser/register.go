package ginuser

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/component/hasher"
	"golang-blog-api/modules/user/userbiz"
	"golang-blog-api/modules/user/usermodel"
	"golang-blog-api/modules/user/userstore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstore.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
