package main

import (
	"fmt"
	"food_delivery/component/appctx"
	restaurantgin "food_delivery/modules/restaurant/transport/gin"
	uploadgin "food_delivery/modules/upload/transport/gin"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello")

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(db, err)

	db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		v1.POST("/upload", uploadgin.UploadImage(appCtx))

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

	r.Run("localhost:3000")

}
