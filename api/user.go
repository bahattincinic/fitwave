package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// getTask godoc
//
//	@Summary	Get Task Detail
//	@Tags		user
//	@Accept		json
//	@Param		id				path		string	true	"Task ID"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	queue.TaskResult
//	@Router		/api/user/task/{id} [get]
func (a *API) getTask(c echo.Context) error {
	task, err := a.q.GetResult(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, task)
}

// runQuery godoc
//
//	@Summary	Run Query
//	@Tags		user
//	@Accept		json
//	@Param		input			body		api.runQuery.queryInput	true	"Query Input"
//	@Param		Authorization	header		string					true	"Bearer <Access Token>"
//	@Success	200				{object}	queue.TaskResult
//	@Failure	400				{object}	ErrorResponse
//	@Router		/api/user/query [post]
func (a *API) runQuery(c echo.Context) error {
	type queryInput struct {
		Query string `json:"query"  validate:"required" err:"query is required"`
	}

	var in queryInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	task := a.q.AddTask(func() (interface{}, error) {
		return a.db.RunQuery(in.Query)
	})

	return c.JSON(http.StatusOK, task)
}
