package model

import (
	"gorm.io/gorm"
)

// Teacher Model
type Teacher struct {
	gorm.Model
	Vorname   string `gorm:"not null" json:"vorname"`
	Nachname  string `gorm:"not null" json:"nachname"`
	CreatedAt string `gorm:"not null" json:"created_at"`
	UserID    int    `gorm:"index" json:"user_id"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}
