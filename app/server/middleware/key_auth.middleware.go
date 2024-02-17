package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ogurilab/school-lunch-api/bootstrap"
)

func KeyAuth(env bootstrap.Env) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(keyAuthConfig(env))
}

func keyAuthConfig(env bootstrap.Env) middleware.KeyAuthConfig {
	return middleware.KeyAuthConfig{
		KeyLookup: "header:X-Admin-Key",
		Validator: func(key string, c echo.Context) (bool, error) {

			return key == env.ADMIN_KEY, nil
		},
	}
}
