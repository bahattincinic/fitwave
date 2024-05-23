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
	requireAuth := a.requireAuthMiddleware()
	api := a.ec.Group("/api")

	api.GET("/", a.serveOK)
	api.GET("/status", a.serveOK)
	api.GET("/docs/*", echoSwagger.WrapHandler)

	{
		gr := api.Group("/gears")
		gr.GET("", a.listGears)
		gr.GET("/:id", a.getGear)
	}

	{
		act := api.Group("/activities")
		act.GET("", a.listActivities)
		act.GET("/:id", a.getActivity, a.setActivityMiddleware())
		act.GET("/:id/gpx", a.exportActivityGPS, a.setActivityMiddleware(), requireAuth)
		act.GET("/:id/laps", a.getActivityLaps, a.setActivityMiddleware(), requireAuth)
	}

	{
		ath := api.Group("/athletes")
		ath.GET("", a.listAthletes)
		ath.GET("/:id", a.getAthlete)
	}

	{
		auth := api.Group("/auth")
		auth.POST("/token", a.getAccessToken)
		auth.GET("/authorization-url", a.getAuthorizationURL)
	}

	{
		usr := api.Group("/user")
		usr.POST("/sync", a.syncData, requireAuth)
		usr.GET("/me", a.getMe, requireAuth)
		usr.GET("/task/:id", a.getTask)
		usr.GET("/config", a.getConfig)
		usr.PUT("/config", a.upsertConfig)
		usr.POST("/query", a.runQuery)
	}

	{
		dash := api.Group("/dashboards")
		// Dashboard
		dash.GET("", a.listDashboards)
		dash.POST("", a.createDashboard)
		dash.GET("/:id", a.getDashboard, a.setDashboardMiddleware())
		dash.PUT("/:id", a.updateDashboard, a.setDashboardMiddleware())
		dash.DELETE("/:id", a.deleteDashboard, a.setDashboardMiddleware())
		dash.POST("/:id/run", a.runQuery, a.setDashboardMiddleware())
		// Component
		dash.GET("/:id/components", a.getDashboardComponents, a.setDashboardMiddleware())
		dash.POST("/:id/components", a.createComponent, a.setDashboardMiddleware())
		dash.PUT("/:id/components/:cpid", a.updateComponent,
			a.setDashboardMiddleware(), a.setComponentMiddleware())
		dash.DELETE("/:id/components/:cpid", a.deleteComponent,
			a.setDashboardMiddleware(), a.setComponentMiddleware())
		dash.POST("/:id/components/:cpid/run", a.runComponent,
			a.setDashboardMiddleware(), a.setComponentMiddleware())
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
	if path == "/" {
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
