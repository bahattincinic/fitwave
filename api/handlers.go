package api

import (
	"fmt"
	"net/http"

	"github.com/bahattincinic/fitwave/api/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *API) setupHandlers() {
	requireAuth := a.requireAuthMiddleware()

	a.ec.GET("/", a.serveOK)
	a.ec.GET("/status", a.serveOK)
	a.ec.GET("/docs/*", echoSwagger.WrapHandler)

	{
		cfg := a.ec.Group("/config")
		cfg.GET("/", a.getConfig)
		cfg.PUT("/", a.upsertConfig)
	}

	{
		gr := a.ec.Group("/gears")
		gr.GET("/", a.listGears)
		gr.GET("/:id", a.getGear)
	}

	{
		act := a.ec.Group("/activities")
		act.GET("/", a.listActivities)
		act.GET("/:id", a.getActivity)
	}

	{
		ath := a.ec.Group("/athletes")
		ath.GET("/", a.listAthletes)
		ath.GET("/:id", a.getAthlete)
	}

	{
		dt := a.ec.Group("/data")
		dt.POST("/sync", a.syncData, requireAuth)
		dt.GET("/task/:id", a.getTask)
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
