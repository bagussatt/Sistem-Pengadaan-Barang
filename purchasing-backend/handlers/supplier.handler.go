package handlers

import (
	"github.com/gofiber/fiber/v2"
	"purchasing-backend/config"
	"purchasing-backend/models"
)


func CreateSupplier(c *fiber.Ctx) error {
	var supplier models.Supplier

	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if supplier.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Supplier name is required",
		})
	}

	if err := config.DB.Create(&supplier).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create supplier",
		})
	}

	return c.JSON(supplier)
}


func GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)
	return c.JSON(suppliers)
}


func GetSupplierByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var supplier models.Supplier
	if err := config.DB.First(&supplier, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Supplier not found",
		})
	}

	return c.JSON(supplier)
}


func UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")

	var supplier models.Supplier
	if err := config.DB.First(&supplier, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Supplier not found",
		})
	}

	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	config.DB.Save(&supplier)

	return c.JSON(supplier)
}


func DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.Supplier{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete supplier",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supplier deleted",
	})
}
