package ginpost

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/post/postbiz"
	"golang-blog-api/modules/post/postmodel"
	"golang-blog-api/modules/post/poststore"
	"golang-blog-api/modules/post/poststore/grpcstore"
	demo "golang-blog-api/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter postmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fullfil()

		store := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		userFavoriteStore := grpcstore.NewGRPCClient(demo.NewFavoriteServiceClient(appCtx.GetGRPCClientConnection()))
		biz := postbiz.NewListBiz(store, userFavoriteStore)

		result, err := biz.List(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
