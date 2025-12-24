package models

import "time"

type Purchasing struct {
	ID         uint               `gorm:"primaryKey"`
	Date       time.Time          `gorm:"not null"`
	SupplierID uint               `gorm:"not null"`
	UserID     uint               `gorm:"not null"`
	GrandTotal float64            `gorm:"type:numeric(14,2)"`
	Details    []PurchasingDetail `gorm:"foreignKey:PurchasingID"`
	CreatedAt  time.Time

	Supplier Supplier `gorm:"foreignKey:SupplierID"`
	User     User     `gorm:"foreignKey:UserID"`
}
