package config

import (
	"challenge-relation-gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	var err error
	dsn := "root:root@tcp(localhost)/heroes_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Hero{}, &models.Meta{})
}

func GetDB() *gorm.DB {
	return DB
}

// heroes_db
