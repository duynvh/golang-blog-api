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

func ListFavoritedPostsOfAUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter favoritemodel.Filter
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		filter.UserId = requester.GetUserId()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fullfil()

		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewListFavoritedPostOfAuthUserBiz(store)

		result, err := biz.ListFavoritedPostsOfAUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
