package validator

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/util"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {

	validator := validator.New()

	validator.RegisterValidation("YYYY-MM-DD", ValidDateFormat)
	validator.RegisterValidation("multipleULID", ValidMultipleULID)

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

func ValidMultipleULID(fl validator.FieldLevel) bool {

	ids, ok := fl.Field().Interface().([]string)
	if !ok {
		return false
	}

	if len(ids) == 0 {
		return true
	}

	limit, ok := fl.Parent().FieldByName("Limit").Interface().(int32)

	if !ok {
		return false
	}

	if len(ids) > int(limit) {
		return false
	}

	for _, id := range ids {
		if _, err := util.ParseUlid(id); err != nil {
			return false
		}
	}

	return true
}
