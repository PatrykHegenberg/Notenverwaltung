package model

import "gorm.io/gorm"

// Address represents a physical address.
type Address struct {
	gorm.Model        // GORM model for database ORM
	Street     string `json:"street"`     // Street name
	Postal     string `json:"postal"`     // Postal code
	City       string `json:"city"`       // City name
	Number     string `json:"number"`     // Building number
	OwnerID    uint   `json:"owner_id"`   // Owner's ID
	OwnerType  string `json:"owner_type"` // Owner's type
}
