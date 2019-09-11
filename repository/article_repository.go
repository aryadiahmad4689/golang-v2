package repository

import (
	"fmt"
	"v2/databases"
	"v2/models"

	"github.com/labstack/echo"
)

//SaveArticle repository
func SaveArticle(c echo.Context) *models.Posts {
	db := databases.DbManager()
	data := new(models.Posts)
	data.Title = c.FormValue("title")
	data.Desc = c.FormValue("Des")
	db.Create(data)
	return data

}

// DataArticle Repository
func DataArticle() []models.Posts {
	db := databases.DbManager()
	data := []models.Posts{}
	db.Find(&data)
	return data
}

// ArticleID ..
func ArticleID(c echo.Context) []models.Posts {
	db := databases.DbManager()
	posts := []models.Posts{}
	id := c.Param("id")
	fmt.Println(id)
	db.Where("id = ?", id).Find(&posts)
	return posts
}

// DeleteByID ..
func DeleteByID(c echo.Context) []models.Posts {
	db := databases.DbManager()
	posts := []models.Posts{}
	id := c.Param("id")
	db.Unscoped().Delete(&posts, "id = ?", id)
	return posts
}

// UpdateByID ..
func UpdateByID(c echo.Context) []models.Posts {
	db := databases.DbManager()
	posts := []models.Posts{}
	id := c.Param("id")
	title := c.FormValue("title")
	desc := c.FormValue("content")
	db.Model(&posts).Where("id= ?", id).Update(map[string]interface{}{"title": title, "desc": desc})
	return posts
}
