package handlers

import "github.com/gofiber/fiber/v2"

func CreatePurchase(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "CreatePurchase endpoint OK",
	})
}
