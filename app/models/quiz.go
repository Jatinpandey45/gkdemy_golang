package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Quiz struct {
	gorm.Model
	Title         string
	Question      string
	QuestionSlug  string
	ScheduleType  string
	Option1       string
	Option2       string
	Option3       string
	Option4       string
	CorrectOption string
	CodeSniff     string
	Explanation   string
	PublishAt     time.Time
	CategoryID    uint
	Category      Category `gorm:"foreignkey:CategoryID"`
}

func (Quiz) TableName() string {
	return "gk_quiz"
}
