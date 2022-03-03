package ginupload

import (
	"golang-blog-api/common"
	"golang-blog-api/component"
	"golang-blog-api/modules/upload/uploadbiz"
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		// db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// imgStore := uploadstore.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
