package restaurantmodel

import (
	"strings"
)

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
	Status  *int    `json:"-" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (u *RestaurantUpdate) Validate() error {
	if strPtr := u.Name; strPtr != nil {
		str := strings.TrimSpace(*strPtr)

		if str == "" {
			return ErrNameCanNotBeBlank
		}

		u.Name = &str
	}

	if strPtr := u.Address; strPtr != nil {
		str := strings.TrimSpace(*strPtr)

		if str == "" {
			return ErrAddressCanNotBeBlank
		}

		u.Address = &str
	}

	return nil
}
