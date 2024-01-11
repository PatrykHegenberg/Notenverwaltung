package model

import (
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Email     string `gorm:"not null;unique" json:"email"`
	Username  string `gorm:"not null;unique" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	CreatedAt string `gorm:"not null" json:"created_at"`
	RoleID    uint   `json:"role_id"`
	Role      Role   `json:"role"`
	SchoolID  uint   `json:"school_id"`
	School    School `json:"school"`
}
