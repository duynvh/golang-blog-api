package ginfavorite

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/favorite/favoritebiz"
	"golang-blog-api/modules/favorite/favoritemodel"
	"golang-blog-api/modules/favorite/favoritestore"
	"golang-blog-api/modules/post/poststore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Favorite(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data favoritemodel.FavoriteCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fakePostId, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()
		data.PostId = int(fakePostId.GetLocalID())

		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		postStore := poststore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewFavoriteBiz(store, postStore)

		if err := biz.Favorite(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
