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

	protected.Post("/items", handlers.CreateItem)
	protected.Get("/items", handlers.GetItems)
	protected.Get("/items/:id", handlers.GetItemByID)
	protected.Put("/items/:id", handlers.UpdateItem)
	protected.Delete("/items/:id", handlers.DeleteItem)

	protected.Post("/suppliers", handlers.CreateSupplier)
	protected.Get("/suppliers", handlers.GetSuppliers)
	protected.Get("/suppliers/:id", handlers.GetSupplierByID)
	protected.Put("/suppliers/:id", handlers.UpdateSupplier)
	protected.Delete("/suppliers/:id", handlers.DeleteSupplier)

}
