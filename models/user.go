package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateOfBirth time.Time

const dateLayout = "2006-01-02"

// UnmarshalJSON parses the date from JSON (no pointer receiver)
func (d *DateOfBirth) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1]) // Remove quotes
	t, err := time.Parse(dateLayout, str)
	if err != nil {
		return fmt.Errorf("could not parse date: %v", err)
	}
	*d = DateOfBirth(t)
	return nil
}

// MarshalJSON serializes the date into JSON (no pointer receiver)
func (d DateOfBirth) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(d).Format(dateLayout))), nil
}

// Value returns the date as a SQL driver value (value receiver)
func (d DateOfBirth) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// Scan converts SQL value into a DateOfBirth (pointer receiver)
func (d *DateOfBirth) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		*d = DateOfBirth(v)
		return nil
	}
	return fmt.Errorf("cannot convert %v to DateOfBirth", value)
}

type User struct {
	ID                     uint        `gorm:"primaryKey"`
	Name                   string      `gorm:"type:varchar(255)"`
	Email                  string      `gorm:"uniqueIndex;type:varchar(255)"`
	Password               string      `gorm:"type:varchar(255)"`
	DateOfBirth            DateOfBirth `gorm:"type:date"`
	Gender                 string      `gorm:"type:varchar(20)"`
	PhoneNumber            string      `gorm:"type:varchar(20)"`
	EducationalInstitution string      `gorm:"type:varchar(255)"`
	Profession             string      `gorm:"type:varchar(255)"`
	Address                string      `gorm:"type:text"`
	Province               string      `gorm:"type:varchar(255)"`
	City                   string      `gorm:"type:varchar(255)"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
