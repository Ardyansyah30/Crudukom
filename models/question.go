package models

import (
	"time"
)

type Question struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	IDPackage     int64     `gorm:"not null"`
	Question      string    `gorm:"type:text;not null"`
	Answer        string    `gorm:"type:varchar(255);default:null"`
	CorrectAnswer string    `gorm:"type:varchar(255);default:null"`
	PacketID      int64     `gorm:"type:bigint(20);default:null"`
	CreatedAt     time.Time `gorm:"type:datetime(3);default:null"`
	UpdatedAt     time.Time `gorm:"type:datetime(3);default:null"`
}
