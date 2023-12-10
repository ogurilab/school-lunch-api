package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/server/validator"
)

func newSetUpTestServer() *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	return e
}
