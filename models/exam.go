package models

import (
	"time"
)

type Exam struct {
	ID            int       `gorm:"primaryKey"`
	IDPacket      int       `gorm:"not null"`
	IDUser        int       `gorm:"not null"`
	NameExam      string    `gorm:"type:varchar(255);not null"`
	StartTime     time.Time `gorm:"not null"`
	EndTime       time.Time `gorm:"not null"`
	PaymentStatus string    `gorm:"type:varchar(50);not null"`
	Score         float64   `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
