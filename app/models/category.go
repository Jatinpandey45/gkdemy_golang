package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	LangID              uint
	CategoryName        string
	CategorySlug        string
	CategoryDescription string
	CategoryIcon        string
	Status              string
}

func (Category) TableName() string {
	return "gk_category"
}
