package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// listGears godoc
//
//	@Summary	List Gears
//	@Tags		gear
//	@Accept		json
//	@Produce	json
//	@Param		limit	query		string	false	"pagination limit"
//	@Param		page	query		string	false	"active page"
//	@Success	200		{object}	PaginatedResponse{Results=[]models.Gear, count=int}
//	@Router		/api/gears [get]
func (a *API) listGears(c echo.Context) error {
	offset, limit, err := a.GetPageAndSize(c, 20)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	count, gears, err := a.db.ListGears(int(offset), int(limit))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, PaginatedResponse{
		Results: gears,
		Count:   count,
	})
}

// getGear godoc
//
//	@Summary	Get Gear
//	@Tags		gear
//	@Accept		json
//	@Param		id	path		string	true	"Gear ID"
//	@Success	200	{object}	models.Gear
//	@Router		/api/gears/{id} [get]
func (a *API) getGear(c echo.Context) error {
	ath, err := a.db.GetGear(c.Param("id"))
	if err != nil {
		return err
	}

	if ath == nil {
		return echo.NewHTTPError(http.StatusNotFound, "gear Not Found")
	}

	return c.JSON(http.StatusOK, ath)
}
