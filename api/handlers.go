package api

import (
	"fmt"
	"net/http"

	"github.com/bahattincinic/fitwave/api/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *API) setupHandlers() {
	requireAuth := a.requireAuthMiddleware()

	a.ec.GET("/", a.serveOK)
	a.ec.GET("/status", a.serveOK)
	a.ec.GET("/docs/*", echoSwagger.WrapHandler)

	{
		gr := a.ec.Group("/gears")
		gr.GET("", a.listGears)
		gr.GET("/:id", a.getGear)
	}

	{
		act := a.ec.Group("/activities")
		act.GET("", a.listActivities)
		act.GET("/:id", a.getActivity, a.setActivityMiddleware())
		act.GET("/:id/gpx", a.exportActivityGPS, a.setActivityMiddleware(), requireAuth)
		act.GET("/:id/laps", a.getActivityLaps, a.setActivityMiddleware(), requireAuth)
	}

	{
		ath := a.ec.Group("/athletes")
		ath.GET("", a.listAthletes)
		ath.GET("/:id", a.getAthlete)
	}

	{
		auth := a.ec.Group("/auth")
		auth.POST("/token", a.getAccessToken)
		auth.GET("/authorization-url", a.getAuthorizationURL)
	}

	{
		usr := a.ec.Group("/user")
		usr.POST("/sync", a.syncData, requireAuth)
		usr.GET("/me", a.getMe, requireAuth)
		usr.GET("/task/:id", a.getTask)
		usr.GET("/config", a.getConfig)
		usr.PUT("/config", a.upsertConfig)
		usr.POST("/query", a.runQuery)
	}

	{
		dash := a.ec.Group("/dashboards")
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
