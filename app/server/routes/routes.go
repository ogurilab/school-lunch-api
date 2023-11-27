package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ogurilab/school-lunch-api/bootstrap"
)

func InitRoutes(env bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	v1 := gin.Group("/v1")

	v1.GET("/health", healthHandler())
}
