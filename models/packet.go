package models

import "time"

type Packet struct {
	ID           uint    `gorm:"primaryKey"`
	NamePacket   string  `gorm:"type:varchar(255)"`
	Description  string  `gorm:"type:text"`
	Price        float64 `gorm:"type:decimal"`
	DurationExam int     `gorm:"type:integer"`
	Payment      string  `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
