package model

import (
	"gorm.io/gorm"
)

// Role Model
type Role struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
}
