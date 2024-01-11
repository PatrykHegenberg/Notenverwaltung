package model

import (
	"gorm.io/gorm"
)

// Grade Model
type Grade struct {
	gorm.Model
	Name int `gorm:"not null" json:"name"`
}
