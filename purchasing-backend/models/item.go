package models

import "time"

type Item struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(150);not null"`
	Stock     int       `gorm:"not null;default:0"`
	Price     float64  `gorm:"type:numeric(12,2);not null"`
	CreatedAt time.Time
}
