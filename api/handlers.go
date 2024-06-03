package api

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bahattincinic/fitwave/api/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *API) setupHandlers() {
	requireStravaAuth := a.requireStravaAuthMiddleware()
	requireAppAuth := a.requireAppAuthMiddleware()
	api := a.ec.Group("/api")

	api.GET("/", a.serveOK)
	api.GET("/status", a.serveOK)
	api.GET("/docs/*", echoSwagger.WrapHandler)

	{
		gr := api.Group("/gears", requireAppAuth)
		gr.GET("", a.listGears)
		gr.GET("/:id", a.getGear)
	}

	{
		str := api.Group("/strava", requireAppAuth)
		str.GET("/activities/:id/gpx", a.exportActivityGPS, a.setActivityMiddleware(), requireStravaAuth)
		str.GET("/activities/:id/laps", a.getActivityLaps, a.setActivityMiddleware(), requireStravaAuth)
		str.POST("/sync", a.syncData, requireStravaAuth)
		str.GET("/me", a.getStravaMe, requireStravaAuth)
		str.POST("/token", a.getStravaAccessToken)
		str.GET("/authorization-url", a.getStravaAuthorizationURL)
	}

	{
		act := api.Group("/activities", requireAppAuth)
		act.GET("", a.listActivities)
		act.GET("/:id", a.getActivity, a.setActivityMiddleware())
	}

	{
		ath := api.Group("/athletes", requireAppAuth)
		ath.GET("", a.listAthletes)
		ath.GET("/:id", a.getAthlete)
	}

	{
		cfg := api.Group("/config")
		cfg.GET("", a.getConfig, requireAppAuth)
		cfg.PUT("", a.upsertConfig, requireAppAuth)
		cfg.GET("/setup", a.checkSetupCompleted)
		cfg.POST("/setup", a.completeSetup)
	}

	{
		auth := api.Group("/auth")
		auth.POST("/token", a.login)
	}

	{
		usr := api.Group("/user", requireAppAuth)
		usr.GET("/task/:id", a.getTask)
		usr.POST("/query", a.runQuery)
	}

	{
		dash := api.Group("/dashboards", requireAppAuth)
		dash.GET("", a.listDashboards)
		dash.POST("", a.createDashboard)
		dash.GET("/:id", a.getDashboard, a.setDashboardMiddleware())
		dash.PUT("/:id", a.updateDashboard, a.setDashboardMiddleware())
		dash.DELETE("/:id", a.deleteDashboard, a.setDashboardMiddleware())
		dash.POST("/:id/run", a.runDashboard, a.setDashboardMiddleware())

	}

	{
		comp := api.Group("/dashboards/:id/components",
			requireAppAuth, a.setDashboardMiddleware())
		comp.GET("", a.getDashboardComponents)
		comp.POST("", a.createComponent)
		comp.PUT("/:cpid", a.updateComponent, a.setComponentMiddleware())
		comp.DELETE("/:cpid", a.deleteComponent, a.setComponentMiddleware())
		comp.POST("/:cpid/run", a.runComponent, a.setComponentMiddleware())
	}

	a.ec.GET("/*", a.serveStatic)
}

func (a *API) serveOK(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (a *API) setupSwagger() {
	docs.SwaggerInfo.Title = "FitWave API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", a.cfg.API.PORT)
}

func (a *API) setupCors() {
	a.ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://127.0.0.1:8080",
		},
	}))
}

func (a *API) serveStatic(c echo.Context) error {
	path := filepath.Clean(c.Request().URL.Path)
	if path == "/" || strings.HasPrefix(path, "/app") {
		path = "index.html"
	}
	path = strings.TrimPrefix(path, "/")

	file, err := uiFS.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	contentType := mime.TypeByExtension(filepath.Ext(path))
	c.Response().Header().Set("Content-Type", contentType)

	stat, err := file.Stat()
	if err == nil && stat.Size() > 0 {
		c.Response().Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	}

	if _, err := io.Copy(c.Response().Writer, file); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error serving file")
	}
	return nil
}
