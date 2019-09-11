package controllers

import (
	"net/http"
	"v2/repository"
	"github.com/labstack/echo"
)

// Get Method
func Get(c echo.Context) error {
	var data interface{}
	data = repository.DataArticle()
	response := repository.Responses(http.StatusOK, "OK", data)
	return c.JSON(http.StatusCreated, response)
}

// GetIDArticle method
func GetIDArticle(c echo.Context) error {
	var data interface{}
	data = repository.ArticleID(c)
	response := repository.Responses(http.StatusOK, "OK", data)
	return c.JSON(http.StatusCreated, response)
}

// Save Method
func Save(c echo.Context) error {
	repository.SaveArticle(c)
	var data interface{}
	response := repository.Responses(http.StatusOK, "Anda Berhasil Menambah Data", data)
	// return c.Redirect(http.StatusMovedPermanently, "/index")
	return c.JSON(http.StatusCreated, response)
}

// DeleteArticle ..
func DeleteArticle(c echo.Context) error {
	var data interface{}
	repository.DeleteByID(c)
	response := repository.Responses(http.StatusOK, "Anda Berhasil Menghapus Data", data)
	return c.JSON(http.StatusCreated, response)

}

// UpdateArticle ..
func UpdateArticle(c echo.Context) error {
	var data interface{}
	repository.UpdateByID(c)
	response := repository.Responses(http.StatusOK, "Anda Berhasil Menambah Data", data)
	return c.JSON(http.StatusCreated, response)

}
