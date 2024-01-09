package validator

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {

	validator := validator.New()

	validator.RegisterValidation("YYYY-MM-DD", ValidDateFormat)

	return &CustomValidator{validator: validator}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func ValidDateFormat(fl validator.FieldLevel) bool {
	date := fl.Field().String()

	_, err := time.Parse("2006-01-02", date)

	return err == nil
}
