package uploadgin

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	uploadbiz "food_delivery/modules/upload/biz"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(err)
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(err)
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(err)
		}

		//imgStore
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider())
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		img.FulFill(appCtx.UploadProvider().GetDomain())

		c.JSON(200, common.SimpleSuccessResponse(img))

	}
}
