package models

import (
	"time"
)

type User struct {
	ID                     uint      `gorm:"primaryKey"`
	Name                   string    `json:"Name"`
	Email                  string    `json:"Email"`
	Password               string    `json:"Password"`
	PhoneNumber            string    `json:"PhoneNumber"`
	DateOfBirth            time.Time `json:"DateOfBirth"`
	Gender                 string    `json:"Gender"`
	EducationalInstitution string    `json:"EducationalInstitution"`
	Profession             string    `json:"Profession"`
	Address                string    `json:"Address"`
	Province               string    `json:"Province"`
	City                   string    `json:"City"`
	CreatedAt              time.Time `json:"CreatedAt"`
	UpdatedAt              time.Time `json:"UpdatedAt"`
}
