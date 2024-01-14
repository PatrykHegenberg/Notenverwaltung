package model

import (
	"gorm.io/gorm"
)

// School Model
type School struct {
	gorm.Model
	Name         string     `gorm:"not null" json:"name"`
	SchoolFormID uint       `gorm:"not null" json:"school_form_id"`
	SchoolForm   SchoolForm `json:"school_form"`
	Address      Address    `gorm:"polymorphic:Owner;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"`
	Users        []User     `json:"users"`
	Classes      []Class    `json:"classes"`
}
