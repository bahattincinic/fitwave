package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	userContextKey = "usr"
)

func (a *API) requireAuthMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken := c.Request().Header.Get("Authorization")
			if accessToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Access token is missing")
			}

			cfg, err := a.db.GetCurrentConfig()
			if err != nil {
				return err
			}

			accessToken = strings.TrimPrefix(accessToken, "Bearer ")
			input, err := a.st.NewUser(cfg, accessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "User is invalid.")
			}

			c.Set(userContextKey, input)
			return next(c)
		}
	}
}
