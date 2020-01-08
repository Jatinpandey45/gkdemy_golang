package models

import (
	"github.com/jinzhu/gorm"
)

type GkLang struct {
	gorm.Model
	LanName  string `gorm:"type:varchar(20);column:lan_name"`
	LangCode string `gorm:"type:varchar(5);column:lang_code"`
	Month    []Month
}

// Set User's table name to be `gk_lang`
func (GkLang) TableName() string {
	return "gk_lang"
}
