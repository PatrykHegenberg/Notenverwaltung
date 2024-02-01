package model

import (
	"gorm.io/gorm"
)

// SchoolForm represents the structure of a school form.
type SchoolForm struct {
	gorm.Model        // gorm.Model provides fields for ID, CreatedAt, UpdatedAt, and DeletedAt
	Name       string `json:"school_form"` // Name is the name of the school form and is tagged for JSON serialization with the key "school_form"
}
