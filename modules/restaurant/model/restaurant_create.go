package restaurantmodel

import (
	"food_delivery/common"
	"strings"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameCanNotBeBlank
	}

	data.Address = strings.TrimSpace(data.Address)

	if data.Address == "" {
		return ErrAddressCanNotBeBlank
	}

	return nil
}
