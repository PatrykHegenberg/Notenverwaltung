package model

import (
	"gorm.io/gorm"
)

// Class Model
type Class struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	IsActive bool   `gorm:"not null" json:"is_active"`
}
