package model

import (
	"gorm.io/gorm"
)

// Score Model
type Score struct {
	gorm.Model
	Points    int    `gorm:"not null" json:"points"`
	Comment   string `gorm:"not null" json:"comment"`
	GradeID   uint   `json:"grade_id"`
	Grade     Grade  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"grade"`
	ExamID    uint   `gorm:"not null" json:"exam_id"`
	StudentID uint   `json:"student_id"`
}
