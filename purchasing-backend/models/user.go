package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
}
