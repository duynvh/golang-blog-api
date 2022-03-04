package ginpost

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/category/categorystore"
	"golang-blog-api/modules/post/postbiz"
	"golang-blog-api/modules/post/postmodel"
	"golang-blog-api/modules/post/poststore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data postmodel.PostUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		categoryStore := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewUpdateBiz(store, categoryStore)

		if err := biz.Update(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
