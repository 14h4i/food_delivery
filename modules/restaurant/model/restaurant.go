package restaurantmodel

import (
	"errors"
	"food_delivery/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name    string         `json:"name" gorm:"column:name;"`
	Address string         `json:"address" gorm:"column:addr;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string { return "restaurants" }

var (
	ErrNameCanNotBeBlank    = errors.New("name can not be blank")
	ErrAddressCanNotBeBlank = errors.New("address can not be blank")
	ErrRestaurantDeleted    = errors.New("restaurant is deleted")
)
