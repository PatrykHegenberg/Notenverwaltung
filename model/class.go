package model

import (
	"gorm.io/gorm"
)

// Class represents a class in a school.
type Class struct {
	gorm.Model           // GORM model for database interaction
	Name       string    `gorm:"not null" json:"name"`       // Name of the class
	IsActive   bool      `gorm:"not null" json:"is_active"`  // Whether the class is active
	Year       string    `gorm:"not null" json:"year"`       // Year of the class
	SchoolID   uint      `gorm:"not null" json:"school_id"`  // ID of the school the class belongs to
	TeacherID  uint      `gorm:"not null" json:"teacher_id"` // ID of the teacher for the class
	Students   []Student `json:"students"`                   // List of students in the class
}
