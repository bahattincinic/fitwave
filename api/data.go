package api

import (
	"net/http"

	"github.com/bahattincinic/fitwave/strava"
	"github.com/labstack/echo/v4"
)

// syncData godoc
//
//	@Summary	Sync Strava data
//	@Tags		data
//	@Accept		json
//	@Param		Authorization	header		string	true	"Strava Access Token"
//	@Success	200				{object}	queue.TaskResult
//	@Router		/data/sync [post]
func (a *API) syncData(c echo.Context) error {
	user := c.Get(userContextKey).(*strava.User)

	task := a.q.AddTask(func() (interface{}, error) {
		if err := a.im.Import(user); err != nil {
			return nil, err
		}
		return nil, nil
	})

	return c.JSON(http.StatusOK, task)
}

// getTask godoc
//
//	@Summary	Get Task Detail
//	@Tags		data
//	@Accept		json
//	@Param		id	path		string	true	"Task ID"
//	@Success	200	{object}	queue.TaskResult
//	@Router		/data/task/{id} [get]
func (a *API) getTask(c echo.Context) error {
	task, err := a.q.GetResult(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, task)
}
