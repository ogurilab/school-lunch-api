package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/repository"
	"github.com/ogurilab/school-lunch-api/server/controller"
	"github.com/ogurilab/school-lunch-api/usecase"
)

func NewMenuWithDishesRouter(group *echo.Group, timeout time.Duration, query db.Query) {

	mr := repository.NewMenuWithDishesRepository(query)
	mc := controller.NewMenuWithDishesController(
		usecase.NewMenuWithDishesUsecase(mr, timeout),
	)

	group.GET("/cities/:code/menus/:id", mc.GetByID)
	group.GET("/cities/:code/menus", mc.FetchByCity)
	group.GET("/menus", mc.Fetch)
}
