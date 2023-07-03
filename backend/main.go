package main

import (
	"memonote/database"
	"memonote/routes/notes"
	"memonote/routes/users"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

type FieldError struct {
	Field   string
	Message string
}

type FieldErrors struct {
	Message string
	Errors  []FieldError
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make([]FieldError, len(validationErrors))
		for _, validationError := range validationErrors {
			errors = append(errors, FieldError{Field: validationError.Field(), Message: "required"})
		}

		return echo.NewHTTPError(http.StatusBadRequest, &FieldErrors{Message: "Validation Errors", Errors: errors})
	}

	return nil
}

func main() {
	godotenv.Load()
	database.CreateConnection()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	notes.RegisterRoutes(e)
	users.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
