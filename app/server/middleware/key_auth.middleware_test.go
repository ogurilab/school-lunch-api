package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/bootstrap"
	"github.com/stretchr/testify/require"
)

func TestKeyAuthMiddleware(t *testing.T) {
	env, err := bootstrap.NewEnv("../../")
	require.NoError(t, err)
	e := echo.New()
	e.Use(KeyAuth(env))

	testCases := []struct {
		name     string
		setUpKey func(t *testing.T, req *http.Request)
		check    func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setUpKey: func(t *testing.T, req *http.Request) {
				req.Header.Set("X-Admin-Key", env.ADMIN_KEY)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Unauthorized",
			setUpKey: func(t *testing.T, req *http.Request) {
				req.Header.Set("X-Admin-Key", "invalid")
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "No Key",
			setUpKey: func(t *testing.T, req *http.Request) {
				// do nothing
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			url := "/test"

			req, err := http.NewRequest(http.MethodGet, url, nil)
			tc.setUpKey(t, req)

			require.NoError(t, err)
			e.GET(url, func(c echo.Context) error {
				return c.String(http.StatusOK, "OK")
			})

			e.ServeHTTP(recorder, req)

			tc.check(t, recorder)
		})
	}
}
