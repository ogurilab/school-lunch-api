package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/repository"
	"github.com/ogurilab/school-lunch-api/server/controller"
	"github.com/ogurilab/school-lunch-api/usecase"
)

func NewAdminRouter(group *echo.Group, timeout time.Duration, query db.Query) {
	mr := repository.NewMenuRepository(query)
	mu := usecase.NewMenuUsecase(mr, timeout)

	dr := repository.NewDishRepository(query)
	du := usecase.NewDishUsecase(dr, timeout)

	ac := controller.NewAdminController(mu, du)

	group.POST("/menus", ac.CreateMenu)
	group.POST("/menus/:id/dishes", ac.CreateDish)
	group.POST("/menus/:id/dishes/bulk", ac.CreateDishes)
}
