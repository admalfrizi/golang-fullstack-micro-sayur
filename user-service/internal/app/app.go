package app

import (
	"user-service/utils/validator"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func RunServer() {
	e := echo.New()
	e.Use(middleware.CORS())

	customValidator := validator.NewValidator()
	en.RegisterDefaultTranslations(customValidator.Validator, customValidator.Translator)
	e.Validator = customValidator
}