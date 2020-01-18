package controllers

import (
	"my-app/app/models"

	_ "github.com/go-sql-driver/mysql"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Category struct {
	gormc.Controller
}

/**
* index page
* @return : application/json
 */

func (c Category) Index() revel.Result {

	var categoryModel []models.Category

	var records = c.DB.Find(&categoryModel)

	return c.RenderJSON(records)
}
