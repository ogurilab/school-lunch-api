package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/bootstrap"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/server/middleware"
)

func InitRoutes(env bootstrap.Env, timeout time.Duration, e *echo.Echo, query db.Query) {

	NewDocumentRouter(e)

	admin := e.Group("/admin")
	admin.Use(middleware.KeyAuth(env))
	NewAdminRouter(admin, timeout, query)

	v1 := e.Group("/v1")

	NewSwaggerRouter(v1)
	NewCityRouter(v1, timeout, query)
	NewMenuRouter(v1, timeout, query)
	NewMenuWithDishesRouter(v1, timeout, query)
	NewDishRouter(v1, timeout, query)
	NewAllergenRouter(v1, timeout, query)

}
