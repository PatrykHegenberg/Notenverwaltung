package model

import (
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	gorm.Model         // Model is a struct that provides basic fields for a GORM model.
	Email      string  `gorm:"not null;unique" json:"email"`                                                             // Email is the user's email address.
	Username   string  `gorm:"not null;unique" json:"username"`                                                          // Username is the user's username.
	Vorname    string  `gorm:"not null;" json:"vorname"`                                                                 // Vorname is the user's first name.
	Nachname   string  `gorm:"not null;" json:"nachname"`                                                                // Nachname is the user's last name.
	Password   string  `gorm:"not null" json:"password"`                                                                 // Password is the user's password.
	IsAdmin    bool    `json:"is_admin"`                                                                                 // IsAdmin indicates whether the user is an administrator.
	SchoolID   uint    `json:"school_id"`                                                                                // SchoolID is the ID of the user's school.
	Address    Address `gorm:"polymorphic:Owner;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"` // Address is the user's address.
}
