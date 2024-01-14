package model

import (
	"gorm.io/gorm"
)

// School Model
type School struct {
	gorm.Model              // Model is a struct that contains the common fields for all models
	Name         string     `gorm:"not null" json:"name"`                                                                     // Name is the name of the school
	SchoolFormID uint       `gorm:"not null" json:"school_form_id"`                                                           // SchoolFormID is the ID of the school form associated with the school
	SchoolForm   SchoolForm `json:"school_form"`                                                                              // SchoolForm is the school form associated with the school
	Address      Address    `gorm:"polymorphic:Owner;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"address"` // Address is the address of the school
	Users        []User     `json:"users"`                                                                                    // Users is a list of users associated with the school
	Classes      []Class    `json:"classes"`                                                                                  // Classes is a list of classes offered by the school
}
