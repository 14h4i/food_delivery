package restaurantgin

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantbiz "food_delivery/modules/restaurant/biz"
	restaurantmodel "food_delivery/modules/restaurant/model"
	restaurantstorage "food_delivery/modules/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&newData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Dependencies install
		u := c.MustGet(common.CurrentUser).(common.Requester)
		newData.OwnerId = u.GetUserId()

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateNewRestaurant(c.Request.Context(), &newData); err != nil {
			panic(err)
		}

		newData.Mask(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newData.FakeId))
	}
}
