package handlers

import (
	"github.com/gofiber/fiber/v2"
	"purchasing-backend/config"
	"purchasing-backend/models"
)

// CreateItem godoc
// @Summary Create new item
// @Description Create a new item with name, stock and price
// @Tags Items
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body models.Item true "Item Request"
// @Success 200 {object} models.Item
// @Failure 400 {object} map[string]string
// @Router /items [post]
func CreateItem(c *fiber.Ctx) error {
	var item models.Item

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if item.Name == "" || item.Price <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name and price are required",
		})
	}

	if err := config.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create item",
		})
	}

	return c.JSON(item)
}

// GetItems godoc
// @Summary Get all items
// @Description Get list of all items
// @Tags Items
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} models.Item
// @Router /items [get]
func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	config.DB.Find(&items)
	return c.JSON(items)
}

// GetItemByID godoc
// @Summary Get item by ID
// @Description Get single item by ID
// @Tags Items
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Failure 404 {object} map[string]string
// @Router /items/{id} [get]
func GetItemByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	return c.JSON(item)
}

// UpdateItem godoc
// @Summary Update item
// @Description Update item by ID
// @Tags Items
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Item ID"
// @Param request body models.Item true "Item Request"
// @Success 200 {object} models.Item
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /items/{id} [put]
func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")

	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Item not found",
		})
	}

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	config.DB.Save(&item)

	return c.JSON(item)
}

// DeleteItem godoc
// @Summary Delete item
// @Description Delete item by ID
// @Tags Items
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Item ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /items/{id} [delete]
func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete item",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item deleted",
	})
}
