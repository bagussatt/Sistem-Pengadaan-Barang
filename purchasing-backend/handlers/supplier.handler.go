package handlers

import (
	"github.com/gofiber/fiber/v2"
	"purchasing-backend/config"
	"purchasing-backend/models"
)

// CreateSupplier godoc
// @Summary Create new supplier
// @Description Create a new supplier with name, email and address
// @Tags Suppliers
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body models.Supplier true "Supplier Request"
// @Success 200 {object} models.Supplier
// @Failure 400 {object} map[string]string
// @Router /suppliers [post]
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

// GetSuppliers godoc
// @Summary Get all suppliers
// @Description Get list of all suppliers
// @Tags Suppliers
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} models.Supplier
// @Router /suppliers [get]
func GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)
	return c.JSON(suppliers)
}

// GetSupplierByID godoc
// @Summary Get supplier by ID
// @Description Get single supplier by ID
// @Tags Suppliers
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.Supplier
// @Failure 404 {object} map[string]string
// @Router /suppliers/{id} [get]
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

// UpdateSupplier godoc
// @Summary Update supplier
// @Description Update supplier by ID
// @Tags Suppliers
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Supplier ID"
// @Param request body models.Supplier true "Supplier Request"
// @Success 200 {object} models.Supplier
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /suppliers/{id} [put]
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

// DeleteSupplier godoc
// @Summary Delete supplier
// @Description Delete supplier by ID
// @Tags Suppliers
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Supplier ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /suppliers/{id} [delete]
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
