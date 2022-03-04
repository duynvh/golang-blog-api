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

func Create(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.PostCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		data.Fulfill()
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		categoryStore := categorystore.NewSQLStore(appCtx.GetMainDBConnection())
		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewCreateBiz(store, categoryStore)

		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
