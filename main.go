package main

import (
	"fmt"
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantgin "food_delivery/modules/restaurant/transport/gin"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	fmt.Println("hello")

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(db, err)

	db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			//Create restaurant
			restaurants.POST("", restaurantgin.CreateRestaurant(appCtx))

			//Get restaurant
			restaurants.GET("/:id", restaurantgin.GetRestaurant(appCtx))

			//Get restaurants
			restaurants.GET("", restaurantgin.ListRestaurant(appCtx))

			//Update restaurant
			restaurants.PUT("/:id", restaurantgin.UpdateRestaurant(appCtx))

			//Delete restaurant
			restaurants.DELETE("/:id", restaurantgin.DeleteRestaurant(appCtx))
		}
	}

	r.Run()

}
