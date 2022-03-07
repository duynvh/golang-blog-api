package ginfavorite

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/favorite/favoritebiz"
	"golang-blog-api/modules/favorite/favoritestore"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unfavorite(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		userId := requester.GetUserId()

		store := favoritestore.NewSQLStore(appCtx.GetMainDBConnection())
		biz := favoritebiz.NewUnfavoriteBiz(store, appCtx.GetPubsub())

		if err := biz.Unfavorite(c.Request.Context(), userId, int(postId.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
