package models

import "time"

type User struct {
	ID                     uint      `gorm:"primaryKey"`              // Perbaikan: tanda kutip ditutup dan menggunakan snake_case
	Name                   string    `json:"name"`                    // Perbaikan: menggunakan snake_case
	Email                  string    `json:"email"`                   // Perbaikan: menggunakan snake_case
	Password               string    `json:"password"`                // Perbaikan: menggunakan snake_case
	PhoneNumber            string    `json:"phone_number"`            // Perbaikan: menggunakan snake_case
	DateOfBirth            time.Time `json:"date_of_birth"`           // Perbaikan: menggunakan snake_case
	Gender                 string    `json:"gender"`                  // Perbaikan: menggunakan snake_case
	EducationalInstitution string    `json:"educational_institution"` // Perbaikan: menggunakan snake_case
	Profession             string    `json:"profession"`              // Perbaikan: menggunakan snake_case
	Address                string    `json:"address"`                 // Perbaikan: menggunakan snake_case
	Province               string    `json:"province"`                // Perbaikan: menggunakan snake_case
	City                   string    `json:"city"`                    // Perbaikan: menggunakan snake_case
	CreatedAt              time.Time `json:"created_at"`              // Perbaikan: menggunakan snake_case
	UpdatedAt              time.Time `json:"updated_at"`              // Perbaikan: menggunakan snake_case
}
