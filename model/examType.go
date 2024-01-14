package model

import (
	"gorm.io/gorm"
)

// ExamType represents the type of an exam.
type ExamType struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"` // Name is the name of the exam type.
}
