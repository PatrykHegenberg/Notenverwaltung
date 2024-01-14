package model

import (
	"gorm.io/gorm"
)

// Score represents a score for a student in an exam.
type Score struct {
	gorm.Model
	Points    int    `gorm:"not null" json:"points"`                                      // Points represents the number of points scored by the student.
	Comment   string `gorm:"not null" json:"comment"`                                     // Comment represents any additional comments about the score.
	GradeID   uint   `json:"grade_id"`                                                    // GradeID represents the ID of the associated grade.
	Grade     Grade  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"grade"` // Grade represents the associated grade.
	ExamID    uint   `gorm:"not null" json:"exam_id"`                                     // ExamID represents the ID of the associated exam.
	StudentID uint   `json:"student_id"`                                                  // StudentID represents the ID of the student.
}
