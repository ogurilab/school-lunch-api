package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/repository"
	"github.com/ogurilab/school-lunch-api/server/controller"
	"github.com/ogurilab/school-lunch-api/usecase"
)

func NewAllergenRouter(group *echo.Group, timeout time.Duration, query db.Query) {
	dr := repository.NewDishRepository(query)
	ar := repository.NewAllergenRepository(query)
	ac := controller.NewAllergenController(
		usecase.NewAllergenUsecase(ar, dr, timeout),
	)

	group.GET("/dishes/:id/allergens", ac.FetchByDishID)
	group.GET("/menus/:id/allergens", ac.FetchByMenuID)
}
