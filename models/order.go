package models

import (
	"time"
)

type Order struct {
	ID            int       `gorm:"primaryKey"`
	IDUser        int       `gorm:"not null"`
	IDPacket      int       `gorm:"not null"`
	PaymentStatus string    `gorm:"type:varchar(255)"`
	OrderDate     time.Time `gorm:"type:date;not null"`
	TotalPrice    float64   `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
