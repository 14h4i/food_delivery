package restaurantstorage

import (
	"context"
	"food_delivery/common"
	restaurantmodel "food_delivery/modules/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db

	var result []restaurantmodel.Restaurant

	if filter.UserId > 0 {
		db = db.Where("owner_id = ?", filter.UserId)
	}

	db = db.Where("status not in (0)")

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Limit(paging.Limit).
		Offset(offset).
		Order("id asc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
