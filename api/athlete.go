package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// listAthletes godoc
//
//	@Summary	List Athletes
//	@Tags		athlete
//	@Accept		json
//	@Produce	json
//	@Param		limit			query		string	false	"pagination limit"
//	@Param		page			query		string	false	"active page"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	PaginatedResponse{Results=[]models.Athlete, count=int}
//	@Router		/api/athletes [get]
func (a *API) listAthletes(c echo.Context) error {
	offset, limit, err := a.GetPageAndSize(c, 20)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	count, athletes, err := a.db.ListAthletes(int(offset), int(limit))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, PaginatedResponse{
		Results: athletes,
		Count:   count,
	})
}

// getActivity godoc
//
//	@Summary	Get Athlete
//	@Tags		athlete
//	@Accept		json
//	@Param		id				path		string	true	"Athlete ID"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	models.Athlete
//	@Router		/api/athletes/{id} [get]
func (a *API) getAthlete(c echo.Context) error {
	ath, err := a.db.GetAthlete(c.Param("id"))
	if err != nil {
		return err
	}

	if ath == nil {
		return echo.NewHTTPError(http.StatusNotFound, "athlete Not Found")
	}

	return c.JSON(http.StatusOK, ath)
}
