package model

import (
	"gorm.io/gorm"
)

// School Model
type School struct {
	gorm.Model
	Name       string `gorm:"not null;unique" json:"name"`
	PLZ        string `gorm:"not null" json:"plz"`
	Street     string `gorm:"not null" json:"street"`
	Homenumber string `gorm:"not null" json:"homenumber"`
}
