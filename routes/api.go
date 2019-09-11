package routes

import (
	"v2/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init Func
func Init() *echo.Echo {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.POST("/store", controllers.Save)
	e.GET("/get", controllers.Get)
	e.GET("/detail/:id", controllers.GetIDArticle)
	e.PUT("/update/:id", controllers.UpdateArticle)
	e.DELETE("/delete/:id", controllers.DeleteArticle)
	return e
}
