package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

/**
 model for post table
**/

type Post struct {
	gorm.Model
	PostTitle     string      `gorm:"type:varchar(250);column:post_title"`
	PostSlug      string      `gorm:"type:text"`
	PostDesc      string      `gorm:"type:text;column:post_desc"`
	Month         Month       `gorm:"foreignkey:MonthID"`
	Category      []*Category `gorm:"many2many:gk_category_post;association_jointable_foreignkey:category_id;jointable_foreignkey:post_id;"`
	MonthID       uint
	LangID        uint
	EmpID         uint
	FeaturedImage string `gorm:"type:text"`
	PublishedAt   *time.Time
	Status        string
	PublishStatus string
	TargetDevice  string
	Tags          []*Tags `gorm:"many2many:gk_tag_post;association_jointable_foreignkey:tag_id;jointable_foreignkey:post_id;"`
}

func (Post) TableName() string {
	return "gk_posts"
}
