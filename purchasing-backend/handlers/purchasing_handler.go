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
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
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

		if err := tx.Model(&purchase).
			Update("grand_total", grandTotal).Error; err != nil {
			return err
		}

		return nil 
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	go utils.SendWebhook(
		os.Getenv("WEBHOOK_URL"),
		fiber.Map{
			"purchase_id": purchase.ID,
			"supplier_id": purchase.SupplierID,
			"user_id":     purchase.UserID,
			"items":       req.Items,
			"grand_total": purchase.GrandTotal,
			"created_at":  time.Now(),
		},
	)

	return c.JSON(fiber.Map{
		"message": "Purchase berhasil",
	})
}
