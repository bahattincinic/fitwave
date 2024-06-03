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

// requireStravaAuthMiddleware validates the presence of a Strava access token in the request header
// and checks if the user is valid. If the token is missing or invalid, it returns an unauthorized error.
func (a *API) requireStravaAuthMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken := c.Request().Header.Get("X-Strava-Authorization")
			if accessToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Strava Access token is missing")
			}

			cfg, err := a.db.GetCurrentConfig()
			if err != nil {
				return err
			}

			input, err := a.st.NewUser(cfg, accessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "User is invalid.")
			}

			c.Set(userContextKey, input)
			return next(c)
		}
	}
}

// requireAppAuthMiddleware validates the presence of an access token in the request header,
// removes the "Bearer " prefix if present, and checks if the token is valid.
// If the token is missing or invalid, it returns an unauthorized error.
func (a *API) requireAppAuthMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accessToken := c.Request().Header.Get("Authorization")
			if accessToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Access token is missing")
			}

			accessToken = strings.TrimPrefix(accessToken, "Bearer ")

			if ok := a.decodeToken(accessToken); !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "User is invalid.")
			}

			return next(c)
		}
	}
}

// setDashboardMiddleware retrieves the dashboard based on the "id" parameter in the request URL,
// sets it in the context, and passes control to the next handler.
// If the dashboard ID is invalid or the dashboard does not exist, it returns a not found error.
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

// setComponentMiddleware retrieves the component based on the "id" and "cpid" parameters in the request URL,
// sets it in the context, and passes control to the next handler.
// If the dashboard ID or component ID is invalid or the component does not exist, it returns a not found error.
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

// setActivityMiddleware retrieves the activity based on the "id" parameter in the request URL,
// sets it in the context, and passes control to the next handler.
// If the activity ID is invalid or the activity does not exist, it returns a not found error.
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
