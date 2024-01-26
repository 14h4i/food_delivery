package usergin

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	userbiz "food_delivery/modules/user/biz"
	usermodel "food_delivery/modules/user/model"
	userstorage "food_delivery/modules/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(common.DbTypeUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))

	}
}
