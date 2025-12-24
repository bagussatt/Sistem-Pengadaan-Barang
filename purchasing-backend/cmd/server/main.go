package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "purchasing-backend/docs" // Swagger docs

	"purchasing-backend/config"
	"purchasing-backend/models"
	"purchasing-backend/routes"
)



func main() {
	_ = godotenv.Load()

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Item{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	routes.Setup(app)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
