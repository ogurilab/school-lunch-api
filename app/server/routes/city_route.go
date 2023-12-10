package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/repository"
	"github.com/ogurilab/school-lunch-api/server/controller"
	"github.com/ogurilab/school-lunch-api/usecase"
)

func NewCityRouter(group *echo.Group, timeout time.Duration, query db.Query) {
	cr := repository.NewCityRepository(query)

	cc := controller.NewCityController(
		usecase.NewCityUsecase(cr, timeout),
	)

	group.GET("/cities/:code", cc.GetByCityCode)

}
