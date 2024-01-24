package main

import (
	"fmt"
	"food_delivery/component/appctx"
	"food_delivery/component/uploadprovider"
	"food_delivery/middleware"
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
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(db, err)

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appctx.NewAppContext(db, s3Provider)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/upload", uploadgin.UploadImage(appCtx))
		v1.GET("/presigned-upload-url", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": s3Provider.GetUploadPresignedURL(c.Request.Context())})
		})

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
