package main

import (
	"clicker/models"
	"clicker/services"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := createDb()
	clickerService := services.NewClickerService(db)
	fmt.Println(clickerService)
}

func createDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("clicker.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.UserClicks{})
	return db
}
