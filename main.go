package main

import (
	"clicker/models"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("clicker.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})
	db.Create(&models.User{Name: "Thiha", Phone: "09993434135"})
	var user models.User
	db.First(&user)
	fmt.Println(user)
}
