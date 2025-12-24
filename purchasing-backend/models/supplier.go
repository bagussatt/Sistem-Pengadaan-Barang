package models

import "time"

type Supplier struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(150);not null"`
	Email     string    `gorm:"type:varchar(150)"`
	Address   string    `gorm:"type:text"`
	CreatedAt time.Time
}
