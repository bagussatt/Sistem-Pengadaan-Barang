package handlers

import (
	"github.com/gofiber/fiber/v2"
	"purchasing-backend/config"
	"purchasing-backend/models"
)


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


func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	config.DB.Find(&items)
	return c.JSON(items)
}


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
