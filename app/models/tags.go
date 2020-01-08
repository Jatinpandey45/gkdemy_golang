package models

import "github.com/jinzhu/gorm"

type Tags struct {
	gorm.Model
	LangID  uint
	TagName string
	TagSlug string
	TagDesc string
}

func (Tags) TableName() string {
	return "gk_tags"
}
