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

// @title			Purchasing Backend API
// @version		1.0
// @description	API untuk Purchasing System Fleetify
// @termsOfService	http://swagger.io/terms/

// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host			localhost:3000
// @BasePath		/api

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token.

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
