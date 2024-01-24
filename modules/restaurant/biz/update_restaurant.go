package restaurantbiz

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/modules/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	Update(
		ctx context.Context,
		cond map[string]interface{},
		updateData *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		}

		return common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, restaurantmodel.ErrRestaurantDeleted)
	}

	if err := biz.store.Update(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
