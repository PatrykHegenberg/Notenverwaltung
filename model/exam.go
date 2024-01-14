package model

import (
	"gorm.io/gorm"
)

// Exam represents an exam.
type Exam struct {
	gorm.Model           // Embedded struct for gorm.Model
	Name        string   `gorm:"not null" json:"name"`        // Name of the exam
	Date        string   `gorm:"not null" json:"date"`        // Date of the exam
	MaxPoints   int      `gorm:"not null" json:"max_points"`  // Maximum points for the exam
	Description string   `gorm:"not null" json:"description"` // Description of the exam
	ExamTypeID  uint     `json:"exam_type_id"`                // ID of the exam type
	ExamType    ExamType `json:"exam_type"`                   // Exam type
	Scores      Score    `json:"scores"`                      // Scores for the exam
}
