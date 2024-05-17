package api

import (
	"fmt"
	"net/http"

	"github.com/bahattincinic/fitwave/api/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *API) setupHandlers() {
	a.ec.GET("/", a.serveOK)
	a.ec.GET("/status", a.serveOK)
	a.ec.GET("/docs/*", echoSwagger.WrapHandler)
}

func (a *API) serveOK(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (a *API) setupSwagger() {
	docs.SwaggerInfo.Title = "FitWave API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", a.cfg.API.PORT)
}
