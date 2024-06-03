package api

import (
	"net/http"

	"github.com/bahattincinic/fitwave/models"
	"github.com/labstack/echo/v4"
)

// getConfig godoc
//
//	@Summary	Get Application Config
//	@Tags		user
//	@Accept		json
//	@Success	200	{object}	models.Config
//	@Router		/api/config [get]
func (a *API) getConfig(c echo.Context) error {
	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cfg)
}

// upsertConfig godoc
//
//	@Summary	Upsert Application Config
//	@Tags		user
//	@Accept		json
//	@Param		config	body		models.Config	true	"Config Input"
//	@Success	200		{object}	models.Config
//	@Router		/api/config [put]
func (a *API) upsertConfig(c echo.Context) error {
	var in models.Config
	if err := c.Bind(&in); err != nil {
		return err
	}

	cfg, err := a.db.UpsertConfig(in)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cfg)
}
