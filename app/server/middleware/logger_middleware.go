package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)

			var event *zerolog.Event

			if err != nil {
				event = log.Error()
				c.Error(err)
			} else {
				event = log.Info()
			}

			end := time.Now()
			latency := end.Sub(start)

			event.
				Str("method", c.Request().Method).
				Str("path", c.Request().URL.Path).
				Int("status", c.Response().Status).
				Dur("latency", latency).
				Msg("request completed")

			return nil
		}
	}
}
