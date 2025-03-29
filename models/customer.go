package models

import (
	"time"
)

type Customer struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Phone     string    `gorm:"not null;unique"`
	Email     string    `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type CustomerInput struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}
