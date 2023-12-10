package server

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/ogurilab/school-lunch-api/bootstrap"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/server/middleware"
	"github.com/ogurilab/school-lunch-api/server/routes"
	"github.com/ogurilab/school-lunch-api/server/validator"
)

func Run(env bootstrap.Env, query db.Query) {
	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	e.Use(middleware.Logger())

	routes.InitRoutes(env, timeout, e, query)

	err := e.Start(env.ServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
