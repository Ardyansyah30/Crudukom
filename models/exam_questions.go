package models

import "time"

type ExamQuestion struct {
	ID         int       `gorm:"primaryKey"`
	IDExam     int       `gorm:"not null"`
	IDQuestion int       `gorm:"not null"`
	UserAnswer int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
