package restaurantmodel

import (
	"errors"
	"food_delivery/common"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

var (
	ErrNameCanNotBeBlank    = errors.New("name can not be blank")
	ErrAddressCanNotBeBlank = errors.New("address can not be blank")
)
