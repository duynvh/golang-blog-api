package gincategory

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/category/categorybiz"
	"golang-blog-api/modules/category/categorymodel"
	"golang-blog-api/modules/category/categorystore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data categorymodel.CategoryUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewUpdateBiz(store)

		if err := biz.Update(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
