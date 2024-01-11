package model

import (
	"gorm.io/gorm"
)

// SubjectExam Model
type SubjectExam struct {
	gorm.Model
	ExamID    int     `gorm:"index" json:"exam_id"`
	SubjectID int     `gorm:"index" json:"subject_id"`
	Exam      Exam    `json:"exam"`
	Subject   Subject `json:"subject"`
}
