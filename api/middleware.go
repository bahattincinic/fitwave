package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	userContextKey      = "usr"
	dashboardContextKey = "dsb"
	componentContextKey = "cp"
	activityContextKey  = "act"
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

func (a *API) setDashboardMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, err := strconv.ParseUint(c.Param("id"), 10, 32)
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprint("invalid dashboard id"))
			}

			dashboard, err := a.db.GetDashboard(uint(id))
			if err != nil {
				return err
			}

			if dashboard == nil {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprint("invalid dashboard id"))
			}

			c.Set(dashboardContextKey, dashboard)
			return next(c)
		}
	}
}

func (a *API) setComponentMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			dshId, err := strconv.ParseUint(c.Param("id"), 10, 32)
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprint("invalid dashboard id"))
			}

			cpId, err := strconv.ParseUint(c.Param("cpid"), 10, 32)
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprint("invalid component id"))
			}

			component, err := a.db.GetComponent(uint(dshId), uint(cpId))
			if err != nil {
				return err
			}

			if component == nil {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprint("invalid component id"))
			}

			c.Set(componentContextKey, component)
			return next(c)
		}
	}
}

func (a *API) setActivityMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			activity, err := a.db.GetActivity(c.Param("id"))
			if err != nil {
				return err
			}
			if activity == nil {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprint("invalid activity id"))
			}

			c.Set(activityContextKey, activity)
			return next(c)
		}
	}
}
