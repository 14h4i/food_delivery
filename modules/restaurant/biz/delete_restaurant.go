package restaurantbiz

import (
	"context"
	"errors"
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
		if err == common.ErrDataNotFound {
			return errors.New("data not found")
		}

		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	zero := 0

	if err := biz.store.Update(ctx,
		map[string]interface{}{"id": id},
		&restaurantmodel.RestaurantUpdate{Status: &zero}); err != nil {
		return err
	}

	return nil
}
