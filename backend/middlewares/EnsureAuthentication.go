package middlewares

import (
	"errors"
	"log"
	"memonote/database"
	"memonote/models"
	"memonote/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EnsureAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			authorizationHeader := c.Request().Header.Get("Authorization")
			if authorizationHeader == "" {
				return c.JSON(http.StatusUnauthorized, utils.APIError{Message: "Unauthorized"})
			}

			data, err := utils.DecodeUserToken(authorizationHeader)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, utils.APIError{Message: "Unauthorized"})
			}

			db := database.GetConnection()
			user := models.User{}

			err = db.Model(&user).Where("id = ?", data.UserID).First(&user).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return c.JSON(http.StatusUnauthorized, utils.APIError{Message: "Unauthorized"})
				} else {
					log.Print(err)
					return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
				}
			}

			c.Set("user", user)
			return next(c)
		}
	}
}
