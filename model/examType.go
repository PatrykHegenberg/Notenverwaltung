package model

import (
	"gorm.io/gorm"
)

// ExamType Model
type ExamType struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
}
