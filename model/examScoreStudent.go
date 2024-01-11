package model

import (
	"gorm.io/gorm"
)

// ExamScoreStudent Model
type ExamScoreStudent struct {
	gorm.Model
	ExamID    int     `gorm:"index" json:"exam_id"`
	ScoreID   int     `gorm:"index" json:"score_id"`
	StudentID int     `gorm:"index" json:"student_id"`
	Exam      Exam    `json:"exam"`
	Score     Score   `json:"score"`
	Student   Student `json:"student"`
}
