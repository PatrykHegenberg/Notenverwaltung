package model

import (
	"gorm.io/gorm"
)

// SubjectTeacherClass Model
type SubjectTeacherClass struct {
	gorm.Model
	ClassID   int     `gorm:"index" json:"class_id"`
	SubjectID int     `gorm:"index" json:"subject_id"`
	TeacherID int     `gorm:"index" json:"teacher_id"`
	Class     Class   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"class"`
	Subject   Subject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subject"`
	Teacher   Teacher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"teacher"`
}
