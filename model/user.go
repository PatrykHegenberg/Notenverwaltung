package model

import (
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Email    string  `gorm:"not null;unique" json:"email"`
	Username string  `gorm:"not null;unique" json:"username"`
	Vorname  string  `gorm:"not null;unique" json:"vorname"`
	Nachname string  `gorm:"not null;unique" json:"nachname"`
	Password string  `gorm:"not null" json:"password"`
	IsAdmin  bool    `json:"is_admin"`
	SchoolID uint    `json:"school_id"`
	Address  Address `gorm:"polymorphic:Owner;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
}
