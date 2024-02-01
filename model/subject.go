package model

import (
	"gorm.io/gorm"
)

// Subject Model
// Subject represents a subject entity in the database.
type Subject struct {
	gorm.Model        // GORM model for managing database records
	Name       string `gorm:"not null" json:"name"` // Name field for the subject, marked as not null in the database and exposed as "name" in JSON
}
