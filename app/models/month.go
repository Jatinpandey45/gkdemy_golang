package models

import (
	"github.com/jinzhu/gorm"
)

type Month struct {
	gorm.Model
	GkLang    GkLang `gorm:"foreignkey:LangID"`
	LangID    uint   `gorm:"column:lang_id"`
	MonthName string `gorm:"type:varchar(50);column:month_name"`
	MonthSlug string `gorm:"type:varchar(50);column:month_slug"`
	MonthDesc string `gorm:"type:text;column:month_desc"`
	Post      []Post
}

func (Month) TableName() string {
	return "gk_month_tags"
}
