package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/ogurilab/school-lunch-api/doc/statik"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog/log"
)

func NewSwaggerRouter(group *echo.Group) {
	statikFS, err := fs.New()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create swagger file system")
	}

	group.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/v1/swagger/", http.FileServer(statikFS))))

}
