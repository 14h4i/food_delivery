package main

import (
	"fmt"
	"food_delivery/common"
	"log"
	"net/http"
	"os"
	"strconv"

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

	/*--------------WRITE--------------*/

	// newRestaurant := Restaurant{Name: "Max", Address: "HHH"}

	// if err := db.Create(&newRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println("Inserted ID:", newRestaurant.Id)

	/*--------------READ--------------*/

	// var oldRestaurant Restaurant

	// if err := db.Table(oldRestaurant.TableName()).Select("id, name").Where("id = ?", 3).First(&oldRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(oldRestaurant)

	// var listRestaurant []Restaurant

	// if err := db.Where("addr = ?", "Somewhere").Limit(10).Find(&listRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// log.Println(listRestaurant)

	// if err := db.Where("addr = ?", "Somewhere").Limit(10).Find(&listRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	/*--------------UPDATE--------------*/

	// dataUpdates := Restaurant{
	// 	Name: "Yan Coffee",
	// }

	// if err := db.Where("id = ?", 3).Limit(10).Updates(&dataUpdates).Error; err != nil {
	// 	log.Println(err)
	// }

	// emptyStr := "ABWQ"

	// dataEmptyUpdates := RestaurantUpdate{
	// 	Name: &emptyStr,
	// }

	// if err := db.Where("id = ?", 4).Limit(10).Updates(&dataEmptyUpdates).Error; err != nil {
	// 	log.Println(err)
	// }

	/*--------------DELETE--------------*/
	// if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 4).Limit(10).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }

	/*--------------REST API--------------*/
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			//Create restaurant
			restaurants.POST("", func(c *gin.Context) {
				var newData RestaurantCreate

				if err := c.ShouldBind(&newData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Create(&newData).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": newData.Id})
			})

			//Get restaurant
			restaurants.GET("/:id", func(c *gin.Context) {
				var data Restaurant

				id, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Where("id = ?", id).First(&data).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": data})
			})

			//Get restaurants
			restaurants.GET("", func(c *gin.Context) {
				var data []Restaurant

				type Paging struct {
					Page  int `json:"page" form:"page"`
					Limit int `json:"limit" form:"limit"`
				}

				var paging Paging

				if err := c.ShouldBind(&paging); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if paging.Page < 1 {
					paging.Page = 1
				}

				if paging.Limit <= 0 {
					paging.Limit = 10
				}

				offset := (paging.Page - 1) * paging.Limit

				if err := db.Offset(offset).Limit(paging.Limit).Order("id asc").Find(&data).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": data})
			})

			//Update restaurant
			restaurants.PUT("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				var data RestaurantUpdate

				if err := c.ShouldBind(&data); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": true})
			})

			//Delete restaurant
			restaurants.DELETE("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": true})
			})
		}
	}

	r.Run()

}
