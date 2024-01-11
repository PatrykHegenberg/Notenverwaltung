package model

import (
	"gorm.io/gorm"
)

// Exam Model
type Exam struct {
	gorm.Model
	Name       string   `gorm:"not null" json:"name"`
	Date       string   `gorm:"not null" json:"date"`
	MaxPoints  int      `gorm:"not null" json:"max_points"`
	Comment    string   `gorm:"not null" json:"comment"`
	ExamTypeID uint     `json:"exam_type_id"`
	ExamType   ExamType `json:"exam_type"`
}
