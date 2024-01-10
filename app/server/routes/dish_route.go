package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/repository"
	"github.com/ogurilab/school-lunch-api/server/controller"
	"github.com/ogurilab/school-lunch-api/usecase"
)

func NewDishRouter(group *echo.Group, timeout time.Duration, query db.Query) {
	dr := repository.NewDishRepository(query)
	dc := controller.NewDishController(usecase.NewDishUsecase(dr, timeout))

	group.GET("/menus/:menuID/dishes", dc.FetchByMenuID)
	group.GET("/dishes/:id", dc.GetByID)
	group.GET("/dishes", dc.Fetch)
}
