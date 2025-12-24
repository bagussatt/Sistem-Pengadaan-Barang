package handlers

import "github.com/gofiber/fiber/v2"

// CreatePurchase godoc
// @Summary Create a new purchase
// @Description Create a new purchase order
// @Tags purchasing
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /purchasing [post]
func CreatePurchase(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "CreatePurchase endpoint OK",
	})
}
