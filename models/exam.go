package models

import (
	"time"
)

type Exam struct {
	ID            int       `gorm:"primaryKey"`                 // Primary key for the exam record
	IDPacket      int       `gorm:"not null"`                   // Foreign key referencing the packet
	IDUser        int       `gorm:"not null"`                   // Foreign key referencing the user
	NameExam      string    `gorm:"type:varchar(255);not null"` // Name of the exam
	Duration      int       `gorm:"not null"`                   // Duration of the exam in minutes
	PaymentStatus string    `gorm:"type:varchar(50);not null"`  // Payment status
	Score         float64   `gorm:"type:decimal(10,2)"`         // Score obtained, with two decimal places
	CreatedAt     time.Time `gorm:"autoCreateTime"`             // Automatically set on creation
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`             // Automatically set on update
}
