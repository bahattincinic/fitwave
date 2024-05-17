package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// listActivities godoc
//
//	@Summary	List Activities
//	@Tags		activity
//	@Accept		json
//	@Produce	json
//	@Param		limit	query		string	false	"pagination limit"
//	@Param		page	query		string	false	"active page"
//	@Success	200		{object}	PaginatedResponse{Results=[]models.Activity, count=int}
//	@Router		/activities/ [get]
func (a *API) listActivities(c echo.Context) error {
	offset, limit, err := a.GetPageAndSize(c, 20)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	count, activities, err := a.db.ListActivities(int(offset), int(limit))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, PaginatedResponse{
		Results: activities,
		Count:   count,
	})
}

// getActivity godoc
//
//	@Summary	Get Activity
//	@Tags		activity
//	@Accept		json
//	@Param		id	path		string	true	"Activity ID"
//	@Success	200	{object}	models.Activity
//	@Router		/activities/{id} [get]
func (a *API) getActivity(c echo.Context) error {
	act, err := a.db.GetActivity(c.Param("id"))
	if err != nil {
		return err
	}

	if act == nil {
		return echo.NewHTTPError(http.StatusNotFound, "activity Not Found")
	}

	return c.JSON(http.StatusOK, act)
}
