package controllers

import (
	"my-app/app/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

/**
* index page
* @return : application/json
 */

func (c App) Index() revel.Result {

	//db, err := gorm.Open("mysql", "root:root@/db_gkdemy?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("mysql", "gkdemy_db_golang:LaFuq*$5_Y&2P_C@tcp(gkdemy.com:9192)/gkdemy_db?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	var posts []models.Post

	var records = db.Preload("Category").Preload("Month").Preload("Month.GkLang").Preload("Tags").Find(&posts)

	return c.RenderJSON(records)
}

func (c App) Store() revel.Result {
	return c.Render()
}
