package controllers

import (
	"errors"
	"log"
	"memonote/database"
	"memonote/dtos"
	"memonote/models"
	"memonote/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct{}

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (uc *UserController) Register(c echo.Context) (err error) {
	credentials := new(dtos.CreateUser)
	if err = c.Bind(credentials); err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	if err = c.Validate(credentials); err != nil {
		return err
	}

	db := database.GetConnection()
	user := models.User{}
	err = db.Model(&user).Where("username = ?", credentials.Username).First(&user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
		}
	} else {
		return c.JSON(http.StatusConflict, utils.APIError{Message: "Username is already taken"})
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 14)

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	hash := string(bytes)

	err = db.Model(&user).Create(&models.User{Username: credentials.Username, Password: hash}).Error
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	token, err := utils.EncodeUserToken(user.ID, user.Username)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	return c.JSON(http.StatusCreated, RegisterResponse{AccessToken: token})
}

func (uc *UserController) Login(c echo.Context) (err error) {
	credentials := new(dtos.Login)
	if err := c.Bind(credentials); err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	if err := c.Validate(credentials); err != nil {
		return err
	}

	db := database.GetConnection()
	user := models.User{}
	err = db.Model(&user).Where("username = ?", credentials.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.APIError{Message: "User not found"})
		} else {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Wrong password"})
	}

	token, err := utils.EncodeUserToken(user.ID, user.Username)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	return c.JSON(http.StatusOK, LoginResponse{AccessToken: token})
}

func (uc *UserController) Me(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, c.Get("user").(models.User))
}
