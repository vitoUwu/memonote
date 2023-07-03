package users

import (
	"memonote/controllers"
	"memonote/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	UserController := new(controllers.UserController)

	router := e.Group("/api/users")
	router.POST("/register", UserController.Register)
	router.POST("/login", UserController.Login)
	router.GET("/me", UserController.Me, middlewares.EnsureAuthentication())
}
