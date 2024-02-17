package controller

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/bootstrap"
	"github.com/ogurilab/school-lunch-api/server/middleware"
	"github.com/ogurilab/school-lunch-api/server/validator"
	"github.com/stretchr/testify/require"
)

func newSetUpTestServer() *echo.Echo {

	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	return e
}

func newSetupAdminTestServer(t *testing.T) (*echo.Echo, bootstrap.Env) {

	env, err := bootstrap.NewEnv("../../")

	require.NoError(t, err)

	e := newSetUpTestServer()
	e.Use(middleware.KeyAuth(env))

	return e, env
}
