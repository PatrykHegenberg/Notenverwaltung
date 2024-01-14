package model

import (
	"gorm.io/gorm"
)

// Student represents a student in the system.
type Student struct {
	gorm.Model // gorm.Model is a struct provided by the GORM library that adds several fields to the struct

	Vorname     string  `gorm:"not null" json:"vorname"`                                                         // Vorname represents the first name of the student
	Nachname    string  `gorm:"not null" json:"nachname"`                                                        // Nachname represents the last name of the student
	DateOfBirth string  `gorm:"not null" json:"date_of_birth"`                                                   // DateOfBirth represents the date of birth of the student
	ClassID     int     `gorm:"index" json:"class_id"`                                                           // ClassID represents the ID of the class the student belongs to
	Address     Address `gorm:"polymorphic:Owner;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"` // Address represents the address of the student
}
