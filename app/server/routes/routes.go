package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/bootstrap"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
)

func InitRoutes(env bootstrap.Env, timeout time.Duration, e *echo.Echo, query db.Query) {
	v1 := e.Group("/v1")

	NewSwaggerRouter(v1)
	NewCityRouter(v1, timeout, query)

}
