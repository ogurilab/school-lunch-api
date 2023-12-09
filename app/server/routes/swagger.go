package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/ogurilab/school-lunch-api/doc/statik"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog/log"
)

func NewSwaggerRouter(group *gin.RouterGroup) {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create swagger file system")
	}

	group.StaticFS("/swagger", statikFS)
}
