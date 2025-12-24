package routes

import (
	"github.com/gofiber/fiber/v2"

	"purchasing-backend/handlers"
	"purchasing-backend/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)

	api.Use(middleware.AuthMiddleware)
	api.Post("/purchasings", handlers.CreatePurchase)
}
