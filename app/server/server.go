package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/ogurilab/school-lunch-api/bootstrap"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/server/routes"
)

func Run(env bootstrap.Env, query db.Query) {
	gin := gin.Default()
	timeout := time.Duration(env.ContextTimeout) * time.Second

	routes.InitRoutes(env, timeout, gin, query)

	err := gin.Run(env.ServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
