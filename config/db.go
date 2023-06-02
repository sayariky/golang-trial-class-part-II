package config

import (
	"fmt"
	"log"
	"trial-class/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := "host=localhost user=postgres password=060818 dbname=trial-class-go port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	} else {
		fmt.Println("connected to db")
		DB = db
	}

	db.AutoMigrate(&entity.Order{}, &entity.Product{})
}
