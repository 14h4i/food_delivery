package userstorage

import (
	"context"
	"food_delivery/common"
	usermodel "food_delivery/modules/user/model"
)

func (s *sqlStore) FindUser(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
