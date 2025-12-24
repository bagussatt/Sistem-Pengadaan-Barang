package routes

import (
	"github.com/gofiber/fiber/v2"

	"purchasing-backend/handlers"
	"purchasing-backend/middleware"
)

func Setup(app *fiber.App) {
	public := app.Group("/api")
	public.Post("/register", handlers.Register)
	public.Post("/login", handlers.Login)

	
	protected := app.Group("/api")
	protected.Use(middleware.AuthMiddleware)

	protected.Post("/purchasings", handlers.CreatePurchase)
	protected.Get("/purchasings", handlers.GetPurchases)
	protected.Get("/purchasings/:id", handlers.GetPurchaseByID)
	protected.Delete("/purchasings/:id", handlers.DeletePurchase)

}
