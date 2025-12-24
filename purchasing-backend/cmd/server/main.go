package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"purchasing-backend/config"
	"purchasing-backend/models"
	"purchasing-backend/routes"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Connect to database
	config.ConnectDB()

	// Auto migrate database schema
	config.DB.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Item{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	)

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.Setup(app)

	// Start server
	app.Listen(":" + os.Getenv("APP_PORT"))
}
