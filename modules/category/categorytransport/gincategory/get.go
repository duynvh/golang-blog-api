package gincategory

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/category/categorybiz"
	"golang-blog-api/modules/category/categorystore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewGetBiz(store)

		data, err := biz.Get(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
