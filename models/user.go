package models

import (
	"fmt"
	"time"
)

type DateOfBirth time.Time

const dateLayout = "2006-01-02"

func (d *DateOfBirth) UnmarshalJSON(b []byte) error {
	str := string(b)
	// Remove the surrounding quotes
	str = str[1 : len(str)-1]
	t, err := time.Parse(dateLayout, str)
	if err != nil {
		return fmt.Errorf("could not parse date: %v", err)
	}
	*d = DateOfBirth(t)
	return nil
}

func (d DateOfBirth) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(d).Format(dateLayout))), nil
}

type User struct {
	ID                             uint      `gorm:"primaryKey"`
	Name                           string    `gorm:"type:varchar(255)"`
	Email                          string    `gorm:"uniqueIndex;type:varchar(255)"`
	Password                       string    `gorm:"type:varchar(255)"`
	DateOfBirth                    time.Time `gorm:"type:date"`
	Gender                         string    `gorm:"type:varchar(20)"`
	PhoneNumber                    string    `gorm:"type:varchar(20)"`
	CurrentEducationalInstitutions string    `gorm:"type:varchar(255)"`
	Profession                     string    `gorm:"type:varchar(255)"`
	Address                        string    `gorm:"type:text"`
	Province                       string    `gorm:"type:varchar(255)"`
	City                           string    `gorm:"type:varchar(255)"`
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
}
