package ginpost

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/post/postbiz"
	"golang-blog-api/modules/post/poststore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewGetBiz(store)

		data, err := biz.Get(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
