package notes

import (
	"memonote/controllers"
	"memonote/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	NoteController := new(controllers.NoteController)

	router := e.Group("/api/notes")
	router.POST("/", NoteController.Create, middlewares.EnsureAuthentication())
	router.GET("/", NoteController.GetAll, middlewares.EnsureAuthentication())
	router.DELETE("/:id", NoteController.Delete, middlewares.EnsureAuthentication())
	router.PATCH("/:id", NoteController.Edit, middlewares.EnsureAuthentication())
	router.GET("/:id", NoteController.GetById, middlewares.EnsureAuthentication())
}
