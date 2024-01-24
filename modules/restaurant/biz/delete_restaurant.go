package restaurantbiz

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/modules/restaurant/model"
)

type DeleteRestaurantStore interface {
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

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
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

	zero := 0

	if err := biz.store.Update(ctx,
		map[string]interface{}{"id": id},
		&restaurantmodel.RestaurantUpdate{Status: &zero}); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
