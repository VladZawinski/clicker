package seeder

import (
	"clicker/middlewares"
	"clicker/models"
	"fmt"

	"gorm.io/gorm"
)

func SeedPredefinedData(db *gorm.DB) {
	result := db.Create(&models.User{
		Name:     "Admin",
		Phone:    "099999999",
		Password: "password",
		Role:     middlewares.AdminRole,
	})
	fmt.Println(result.RowsAffected)
}
