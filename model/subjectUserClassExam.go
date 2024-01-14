package model

import (
	"gorm.io/gorm"
)

// SubjectUserClassExam represents the relationship between a user, subject, class, and exam.
type SubjectUserClassExam struct {
	gorm.Model
	UserID    uint    `json:"user_id"`    // UserID is the ID of the user.
	User      User    `json:"user"`       // User is the user associated with this record.
	SubjectID uint    `json:"subject_id"` // SubjectID is the ID of the subject.
	Subject   Subject `json:"subject"`    // Subject is the subject associated with this record.
	ClassID   uint    `json:"class_id"`   // ClassID is the ID of the class.
	Class     Class   `json:"class"`      // Class is the class associated with this record.
	ExamID    uint    `json:"exam_id"`    // ExamID is the ID of the exam.
	Exam      Exam    `json:"exam"`       // Exam is the exam associated with this record.
}
