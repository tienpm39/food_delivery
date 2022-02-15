package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
}

func (Restaurant) TableName() string {
	return "restaurant"
}

func main() {
	dsn := os.Getenv("DSN")
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

	// create database
	newRestaurant := Restaurant{
		Name: "Phuc Loc Tho",
		Addr: "Nguyen Anh Thu",
	}

	if err := db.Create(&newRestaurant); err != nil {
		fmt.Println(err)
	}

	//select 1 list
	var restaurants []Restaurant
	db.Where("status = ?", 1).Find(&restaurants)
	fmt.Println(restaurants)

	//select 1 dong
	var restaurant Restaurant
	if err := db.Where("id=1").First(&restaurant); err != nil {
		log.Println(err)
	}

	//delete
	db.Table(Restaurant{}.TableName()).Where("id=?", 1).Delete(nil)

	//Update
	newName := "Com Chay Thien Nhan"
	db.Table(Restaurant{}.TableName()).Where("id=?", 1).Updates(&RestaurantUpdate{&newName})

	//Lenh SQL
	//db.Exec("")
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r.Run()
}

//CREATE TABLE `restaurants` (
//	`id` int(11) NOT NULL AUTO_INCREMENT,
//	`owner_id` int(11) NOT NULL,
//	`name` varchar(50) NOT NULL,
//	`addr` varchar(255) NOT NULL,
//	`city_id` int(11) DEFAULT NULL,
//	`lat` double DEFAULT NULL,
//	`lng` double DEFAULT NULL,
//	`cover` json NOT NULL,
//	`logo` json NOT NULL,
//	`shipping_fee_per_km` double DEFAULT '0',
//	`status` int(11) NOT NULL DEFAULT '1',
//	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//	PRIMARY KEY (`id`),
//	KEY `owner_id` (`owner_id`) USING BTREE,
//	KEY `city_id` (`city_id`) USING BTREE,
//	KEY `status` (`status`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
