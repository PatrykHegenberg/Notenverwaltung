package model

import (
	"gorm.io/gorm"
)

// Class Model
type Class struct {
	gorm.Model
	Name      string    `gorm:"not null" json:"name"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	Year      string    `gorm:"not null" json:"year"`
	SchoolID  uint      `gorm:"not null" json:"school_id"`
	TeacherID uint      `gorm:"not null" json:"teacher_id"`
	Students  []Student `json:"students"`
}
