package controllers

import (
	"errors"
	"log"
	"memonote/database"
	"memonote/dtos"
	"memonote/models"
	"memonote/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type NoteController struct{}

func (nc *NoteController) Create(c echo.Context) (err error) {
	requestBody := new(dtos.CreateNote)
	if err := c.Bind(requestBody); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	if err := c.Validate(requestBody); err != nil {
		return err
	}

	db := database.GetConnection()
	note := models.Note{Content: requestBody.Content, Title: requestBody.Title, UserID: c.Get("user").(models.User).ID}
	err = db.Create(&note).Error
	if err != nil {
		log.Fatal(err)
		return
	}

	return c.JSON(http.StatusCreated, note)
}

func (nc *NoteController) Delete(c echo.Context) (err error) {
	noteId := c.Param("id")
	_, err = uuid.Parse(noteId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid note id"})
	}

	db := database.GetConnection()
	note := models.Note{}
	err = db.Model(&note).Where("id = ?", noteId).Where("user_id = ?", c.Get("user").(models.User).ID).First(&note).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.APIError{Message: "Note not found"})
		}
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	err = db.Model(&note).Where("id = ?", noteId).Delete(&note).Error
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	return c.NoContent(http.StatusOK)
}

func (nc *NoteController) Edit(c echo.Context) (err error) {
	noteId := c.Param("id")
	_, err = uuid.Parse(noteId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid note id"})
	}

	body := new(dtos.EditNoteContent)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid request body"})
	}

	if err := c.Validate(body); err != nil {
		return err
	}

	db := database.GetConnection()
	note := models.Note{}
	err = db.Model(&note).Where("id = ?", noteId).Where("user_id = ?", c.Get("user").(models.User).ID).First(&note).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.APIError{Message: "Note not found"})
		}
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	if body.Content != nil {
		note.Content = *body.Content
	}

	if body.Title != nil {
		note.Title = *body.Title
	}

	err = db.Save(&note).Error
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	return c.JSON(http.StatusOK, note)
}

func (nc *NoteController) GetAll(c echo.Context) (err error) {
	db := database.GetConnection()
	notes := []models.Note{}
	err = db.Model(&models.Note{}).Where("user_id = ?", c.Get("user").(models.User).ID).Find(&notes).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.APIError{Message: "No notes found"})
		}
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	return c.JSON(http.StatusOK, notes)
}

func (nc *NoteController) GetById(c echo.Context) (err error) {
	noteId := c.Param("id")
	_, err = uuid.Parse(noteId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.APIError{Message: "Invalid note id"})
	}

	db := database.GetConnection()
	note := models.Note{}
	err = db.Model(&note).Where("id = ?", noteId).Where("user_id = ?", c.Get("user").(models.User).ID).First(&note).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, utils.APIError{Message: "Note not found"})
		}
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, utils.APIError{Message: "An error has occurred"})
	}

	return c.JSON(http.StatusOK, note)
}
