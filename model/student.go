package model

import (
	"gorm.io/gorm"
)

// Student Model
type Student struct {
	gorm.Model
	Vorname   string `gorm:"not null" json:"vorname"`
	Nachname  string `gorm:"not null" json:"nachname"`
	GebDatum  string `gorm:"not null" json:"geb_datum"`
	CreatedAt string `gorm:"not null" json:"created_at"`
	ClassID   int    `gorm:"index" json:"class_id"`
	UserID    int    `gorm:"index" json:"user_id"`
	Class     Class  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"class"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}
