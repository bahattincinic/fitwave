package api

import (
	"net/http"

	"github.com/bahattincinic/fitwave/strava"
	"github.com/labstack/echo/v4"
)

// syncData godoc
//
//	@Summary	Sync Strava data
//	@Tags		user
//	@Accept		json
//	@Param		Authorization	header		string	true	"Strava Access Token"
//	@Success	200				{object}	queue.TaskResult
//	@Router		/user/sync [post]
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
//	@Tags		user
//	@Accept		json
//	@Param		id	path		string	true	"Task ID"
//	@Success	200	{object}	queue.TaskResult
//	@Router		/user/task/{id} [get]
func (a *API) getTask(c echo.Context) error {
	task, err := a.q.GetResult(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, task)
}

// getMe godoc
//
//	@Summary	Get Current User Details
//	@Tags		user
//	@Accept		json
//	@Success	200	{object}	strava.User
//	@Router		/user/me [get]
func (a *API) getMe(c echo.Context) error {
	user := c.Get(userContextKey).(*strava.User)
	return c.JSON(http.StatusOK, user)
}

// runQuery godoc
//
//	@Summary	Run Query
//	@Tags		user
//	@Accept		json
//	@Param		input	body		api.runQuery.queryInput	true	"Query Input"
//	@Success	200		{object}	queue.TaskResult
//	@Router		/user/query [post]
func (a *API) runQuery(c echo.Context) error {
	type queryInput struct {
		Query string `json:"query"`
	}

	var in queryInput
	if err := c.Bind(&in); err != nil {
		return err
	}

	task := a.q.AddTask(func() (interface{}, error) {
		return a.db.RunQuery(in.Query)
	})

	return c.JSON(http.StatusOK, task)
}
