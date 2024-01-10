package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/repository"
	"github.com/ogurilab/school-lunch-api/server/controller"
	"github.com/ogurilab/school-lunch-api/usecase"
)

func NewMenuRouter(group *echo.Group, timeout time.Duration, query db.Query) {
	newMenuBasicRouter(group, timeout, query)
	newMenuWithDishesRouter(group, timeout, query)
}

func newMenuBasicRouter(
	group *echo.Group,
	timeout time.Duration,
	query db.Query,
) {
	mr := repository.NewMenuRepository(query)
	mc := controller.NewMenuController(usecase.NewMenuUsecase(mr, timeout))

	group.GET("/cities/:code/menus/:id/basic", mc.GetByID)
	group.GET("/cities/:code/menus/basic", mc.FetchByCity)

	group.POST("/cities/:code/menus", mc.Create)

}

func newMenuWithDishesRouter(
	group *echo.Group,
	timeout time.Duration,
	query db.Query,
) {
	mr := repository.NewMenuWithDishesRepository(query)
	mc := controller.NewMenuWithDishesController(
		usecase.NewMenuWithDishesUsecase(mr, timeout),
	)

	group.GET("/cities/:code/menus/:id", mc.GetByID)
	group.GET("/cities/:code/menus", mc.FetchByCity)
	group.GET("/menus", mc.Fetch)
}
