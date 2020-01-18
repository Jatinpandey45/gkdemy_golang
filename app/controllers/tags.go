package controllers

import (
	"my-app/app/models"

	_ "github.com/go-sql-driver/mysql"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Tags struct {
	gormc.Controller
}

/**
* index page
* @return : application/json
 */

func (c Tags) Index() revel.Result {

	var tagModel []models.Tags

	var records = c.DB.Find(&tagModel)

	return c.RenderJSON(records)
}
