package domain

import "github.com/labstack/echo/v4"

type AdminController interface {
	CreateMenu(c echo.Context) error
}
