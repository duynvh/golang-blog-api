package ginfavorite

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/favorite/favoritebiz"
	"golang-blog-api/modules/favorite/favoritemodel"
	"golang-blog-api/modules/favorite/favoritestore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsersFavoritedAPost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter favoritemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.Fulfill()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fullfil()

		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewListUsersFavoritedAPostBiz(store)

		result, err := biz.ListUserFavoritedAPost(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
