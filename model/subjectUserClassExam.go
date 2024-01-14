package model

import (
	"gorm.io/gorm"
)

type SubjectUserClassExam struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	User      User    `json:"user"`
	SubjectID uint    `json:"subject_id"`
	Subject   Subject `json:"subject"`
	ClassID   uint    `json:"class_id"`
	Class     Class   `json:"class"`
	ExamID    uint    `json:"exam_id"`
	Exam      Exam    `json:"exam"`
}
