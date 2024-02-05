package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/ogurilab/school-lunch-api/doc/statiks/document"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog/log"
)

func NewDocumentRouter(e *echo.Echo) {
	fs, err := fs.NewWithNamespace("document")

	if err != nil {
		log.Fatal().Err(err).Msg("failed to create document file system")
	}

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", http.FileServer(fs))))
}
