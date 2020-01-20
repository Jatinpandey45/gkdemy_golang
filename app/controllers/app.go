package controllers

import (
	"fmt"
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

	pageLimit := 20

	pageOffset := 0

	c.Params.Query = c.Request.URL.Query()

	var page int

	var count int64

	c.Params.Bind(&page, "page")

	pageOffset = (page - 1) * pageLimit

	fmt.Println(pageOffset)

	var returnResult = make(map[string]interface{})

	records := c.DB.Preload("Category").Preload("Month").Preload("Month.GkLang").Preload("Tags").Order("created_at desc").Limit(pageLimit).Offset(pageOffset).Find(&posts)
	c.DB.Table("gk_posts").Count(&count)

	returnResult["data"] = records.Value
	returnResult["total"] = count

	if page == 0 {
		page = 1
	}

	returnResult["page"] = page

	return c.RenderJSON(returnResult)
}

func (c App) Quiz() revel.Result {

	var quiz []models.Quiz

	pageLimit := 20

	pageOffset := 0

	c.Params.Query = c.Request.URL.Query()

	var page int

	var count int64

	c.Params.Bind(&page, "page")

	pageOffset = (page - 1) * pageLimit

	fmt.Println(pageOffset)

	var returnResult = make(map[string]interface{})

	var records = c.DB.Preload("Category").Limit(pageLimit).Offset(pageOffset).Find(&quiz)

	c.DB.Table("gk_quiz").Count(&count)

	returnResult["data"] = records.Value
	returnResult["total"] = count

	if page == 0 {
		page = 1
	}

	returnResult["page"] = page

	return c.RenderJSON(returnResult)
}
