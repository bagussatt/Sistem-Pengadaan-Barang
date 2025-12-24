package handlers

import (
	"os"
	"time"

	"purchasing-backend/config"
	"purchasing-backend/models"
	"purchasing-backend/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PurchaseItemReq struct {
	ItemID uint `json:"item_id"`
	Qty    int  `json:"qty"`
}

type PurchaseReq struct {
	SupplierID uint              `json:"supplier_id"`
	Items      []PurchaseItemReq `json:"items"`
}

func CreatePurchase(c *fiber.Ctx) error {
	var req PurchaseReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	userID := c.Locals("user_id").(uint)
	var purchase models.Purchasing

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var grandTotal float64

		purchase = models.Purchasing{
			Date:       time.Now(),
			SupplierID: req.SupplierID,
			UserID:     userID,
		}

		if err := tx.Create(&purchase).Error; err != nil {
			return err
		}

		for _, item := range req.Items {
			var dbItem models.Item
			if err := tx.First(&dbItem, item.ItemID).Error; err != nil {
				return err
			}

			if dbItem.Stock < item.Qty {
				return fiber.NewError(400, "Stok tidak cukup")
			}

			subTotal := float64(item.Qty) * dbItem.Price
			grandTotal += subTotal

			if err := tx.Create(&models.PurchasingDetail{
				PurchasingID: purchase.ID,
				ItemID:       item.ItemID,
				Qty:          item.Qty,
				SubTotal:     subTotal,
			}).Error; err != nil {
				return err
			}

			if err := tx.Model(&dbItem).
				Update("stock", dbItem.Stock-item.Qty).Error; err != nil {
				return err
			}
		}

		return tx.Model(&purchase).
			Update("grand_total", grandTotal).Error
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	go utils.SendWebhook(os.Getenv("WEBHOOK_URL"), purchase)

	return c.JSON(fiber.Map{"message": "Purchase berhasil"})
}
func GetPurchases(c *fiber.Ctx) error {
	var purchases []models.Purchasing

	config.DB.
		Preload("Supplier").
		Preload("User").
		Find(&purchases)

	return c.JSON(purchases)
}
func GetPurchaseByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var purchase models.Purchasing
	if err := config.DB.
		Preload("Supplier").
		Preload("User").
		Preload("Details.Item").
		First(&purchase, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Data tidak ditemukan"})
	}

	return c.JSON(purchase)
}
func DeletePurchase(c *fiber.Ctx) error {
	id := c.Params("id")

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var details []models.PurchasingDetail
		tx.Where("purchasing_id = ?", id).Find(&details)

		for _, d := range details {
			tx.Model(&models.Item{}).
				Where("id = ?", d.ItemID).
				Update("stock", gorm.Expr("stock + ?", d.Qty))
		}

		tx.Where("purchasing_id = ?", id).
			Delete(&models.PurchasingDetail{})

		return tx.Delete(&models.Purchasing{}, id).Error
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Purchase dihapus"})
}
