package main

import (
	"clicker/handlers"
	"clicker/models"
	"clicker/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := createDb()
	app := fiber.New()
	// seeder.SeedPredefinedData(db)
	clickerService := services.NewClickerService(db)
	handlers.SetUpHandlers(app, &clickerService)
	app.Listen(":3000")
}

func createDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("clicker.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.UserClicks{})
	return db
}
