package model

import (
	"gorm.io/gorm"
)

type SchoolForm struct {
	gorm.Model
	Name string `json:"school_form"`
}
