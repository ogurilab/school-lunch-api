package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/ogurilab/school-lunch-api/bootstrap"
	"github.com/ogurilab/school-lunch-api/server/routes"
)

func Run(env bootstrap.Env) {
	gin := gin.Default()
	timeout := time.Duration(env.ContextTimeout) * time.Second

	routes.InitRoutes(env, timeout, gin)

	err := gin.Run(env.ServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
