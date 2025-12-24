package models

import "time"

type PurchasingDetail struct {
	ID           uint      `gorm:"primaryKey"`
	PurchasingID uint      `gorm:"not null"`
	ItemID       uint      `gorm:"not null"`
	Qty          int       `gorm:"not null"`
	SubTotal     float64   `gorm:"type:numeric(14,2);not null"`
	CreatedAt    time.Time

	Item Item `gorm:"foreignKey:ItemID"`
}
	