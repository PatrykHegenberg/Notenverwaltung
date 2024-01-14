package model

import (
	"gorm.io/gorm"
)

// Student Model
type Student struct {
	gorm.Model
	Vorname     string  `gorm:"not null" json:"vorname"`
	Nachname    string  `gorm:"not null" json:"nachname"`
	DateOfBirth string  `gorm:"not null" json:"date_of_birth"`
	ClassID     int     `gorm:"index" json:"class_id"`
	Address     Address `gorm:"polymorphic:Owner;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
}
