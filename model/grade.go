package model

import (
	"gorm.io/gorm"
)

// Grade represents a grade with a name.
type Grade struct {
	gorm.Model
	Name int `gorm:"not null" json:"name"` // Name is the name of the grade.
}
