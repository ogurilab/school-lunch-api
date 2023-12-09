package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ogurilab/school-lunch-api/bootstrap"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
)

func InitRoutes(env bootstrap.Env, timeout time.Duration, gin *gin.Engine, query db.Query) {
	v1 := gin.Group("/v1")

	NewSwaggerRouter(v1)

}
