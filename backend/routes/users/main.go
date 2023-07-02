package users

import (
	"memonote/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	UserController := new(controllers.UserController)

	router := e.Group("/users")
	router.POST("/register", UserController.Register)
	router.POST("/login", UserController.Login)
}
