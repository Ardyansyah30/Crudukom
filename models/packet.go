package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Packet struct {
	ID           uint      `gorm:"primaryKey"`         // Primary key for the packet
	NamePacket   string    `gorm:"type:varchar(255)"`  // Name of the packet
	Description  string    `gorm:"type:text"`          // Description of the packet
	Price        float64   `gorm:"type:decimal(10,2)"` // Price of the packet with precision (e.g., 10.2 for monetary values)
	DurationExam int       `gorm:"type:integer"`       // Duration of the exam associated with the packet
	Payment      string    `gorm:"type:varchar(255)"`  // Payment method or status
	CreatedAt    time.Time `gorm:"autoCreateTime"`     // Automatically set the creation time
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`     // Automatically set the update time
}

// Custom UnmarshalJSON for the Price field to handle string or number input
func (p *Packet) UnmarshalJSON(data []byte) error {
	// Temporary structure to hold incoming data
	var tmp struct {
		ID           uint        `json:"id"`
		NamePacket   string      `json:"name_packet"`
		Description  string      `json:"description"`
		Price        interface{} `json:"price"` // Accept interface{} so we can handle both string and number
		DurationExam int         `json:"duration_exam"`
		Payment      string      `json:"payment"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
	}

	// Unmarshal the incoming JSON into the temporary struct
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	// Assign the fields
	p.ID = tmp.ID
	p.NamePacket = tmp.NamePacket
	p.Description = tmp.Description
	p.DurationExam = tmp.DurationExam
	p.Payment = tmp.Payment
	p.CreatedAt = tmp.CreatedAt
	p.UpdatedAt = tmp.UpdatedAt

	// Handle Price, which can be either a string or a number
	switch value := tmp.Price.(type) {
	case string:
		parsedPrice, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("could not parse price string: %v", err)
		}
		p.Price = parsedPrice
	case float64:
		p.Price = value
	default:
		return fmt.Errorf("invalid type for price field")
	}

	return nil
}
