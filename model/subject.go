package model

import (
	"gorm.io/gorm"
)

// Subject Model
type Subject struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}
