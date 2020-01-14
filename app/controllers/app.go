package controllers

import (
	"my-app/app/models"
	
	_ "github.com/go-sql-driver/mysql"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type App struct {
	gormc.Controller
}

/**
* index page
* @return : application/json
 */

func (c App) Index() revel.Result {

	var posts []models.Post

	var records = c.DB.Preload("Category").Preload("Month").Preload("Month.GkLang").Preload("Tags").Find(&posts)

	return c.RenderJSON(records)
}

func (c App) Quiz() revel.Result {

	var quiz []models.Quiz

	var records = c.DB.Preload("Category").Find(&quiz)

	return c.RenderJSON(records)
}
