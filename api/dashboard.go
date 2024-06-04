package api

import (
	"encoding/json"
	"net/http"

	"github.com/bahattincinic/fitwave/models"
	"github.com/labstack/echo/v4"
)

// listDashboards godoc
//
//	@Summary	List Dashboards
//	@Tags		dashboard
//	@Accept		json
//	@Produce	json
//	@Param		limit			query		string	false	"pagination limit"
//	@Param		page			query		string	false	"active page"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	PaginatedResponse{Results=[]models.Dashboard, count=int}
//	@Router		/api/dashboards [get]
func (a *API) listDashboards(c echo.Context) error {
	offset, limit, err := a.GetPageAndSize(c, 20)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	count, dashboards, err := a.db.ListDashboards(int(offset), int(limit))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, PaginatedResponse{
		Results: dashboards,
		Count:   count,
	})
}

// createDashboard godoc
//
//	@Summary	Create Dashboard
//	@Tags		dashboard
//	@Accept		json
//	@Produce	json
//	@Param		input			body		api.createDashboard.dashboardInput	true	"Dashboard Input"
//	@Param		Authorization	header		string								true	"Bearer <Access Token>"
//	@Success	201				{object}	models.Dashboard
//	@Failure	400				{object}	ErrorResponse
//	@Router		/api/dashboards [post]
func (a *API) createDashboard(c echo.Context) error {
	type dashboardInput struct {
		Name string `json:"name" validate:"required" err:"name is required"`
	}

	var in dashboardInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	dashboard, err := a.db.CreateDashboard(in.Name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dashboard)
}

// getDashboard godoc
//
//	@Summary	Get Dashboard
//	@Tags		dashboard
//	@Accept		json
//	@Param		id				path		string	true	"Dashboard ID"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	models.Dashboard
//	@Router		/api/dashboards/{id} [get]
func (a *API) getDashboard(c echo.Context) error {
	dashboard := c.Get(dashboardContextKey).(*models.Dashboard)

	return c.JSON(http.StatusOK, dashboard)
}

// updateDashboard godoc
//
//	@Summary	Update Dashboard
//	@Tags		dashboard
//	@Accept		json
//	@Param		id				path		string	true	"Dashboard ID"
//	@Success	200				{object}	models.Dashboard
//	@Param		input			body		api.updateDashboard.dashboardInput	true	"Dashboard Input"
//	@Param		Authorization	header		string								true	"Bearer <Access Token>"
//	@Failure	400				{object}	ErrorResponse
//	@Router		/api/dashboards/{id} [put]
func (a *API) updateDashboard(c echo.Context) error {
	dashboard := c.Get(dashboardContextKey).(*models.Dashboard)

	type dashboardInput struct {
		Name string `json:"name" validate:"required" err:"name is required"`
	}

	var in dashboardInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	dashboard.Name = in.Name
	if err := a.db.UpdateDashboard(dashboard); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dashboard)
}

// deleteDashboard godoc
//
//	@Summary	Delete Dashboard
//	@Tags		dashboard
//	@Accept		json
//	@Param		id				path	string	true	"Dashboard ID"
//	@Param		Authorization	header	string	true	"Bearer <Access Token>"
//	@Success	204
//	@Router		/api/dashboards/{id} [delete]
func (a *API) deleteDashboard(c echo.Context) error {
	dashboard := c.Get(dashboardContextKey).(*models.Dashboard)

	if err := a.db.DeleteDashboard(dashboard); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// runDashboard godoc
//
//	@Summary	Run Dashboard
//	@Tags		dashboard
//	@Accept		json
//	@Param		id				path		string	true	"Dashboard ID"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	queue.TaskResult
//	@Router		/api/dashboards/{id}/run [post]
func (a *API) runDashboard(c echo.Context) error {
	dashboard := c.Get(dashboardContextKey).(*models.Dashboard)

	components, err := a.db.ListComponents(dashboard.ID)
	if err != nil {
		return err
	}

	type componentResult struct {
		ID      uint        `json:"id"`
		Results interface{} `json:"results"`
	}

	task := a.q.AddTask(func() (interface{}, error) {
		var results []componentResult

		for _, com := range components {
			resp, err := a.db.RunQuery(com.Query)
			if err != nil {
				return nil, err
			}
			results = append(results, componentResult{
				ID:      com.ID,
				Results: resp,
			})
		}

		return results, nil
	})

	return c.JSON(http.StatusOK, task)
}

// getDashboardComponents godoc
//
//	@Summary	List Dashboard Components
//	@Tags		dashboard
//	@Accept		json
//	@Produce	json
//	@Param		id				path		string	true	"Dashboard ID"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	PaginatedResponse{Results=[]models.Component, count=int}
//	@Router		/api/dashboards/{id}/components [get]
func (a *API) getDashboardComponents(c echo.Context) error {
	dashboard := c.Get(dashboardContextKey).(*models.Dashboard)

	components, err := a.db.ListComponents(dashboard.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, PaginatedResponse{
		Results: components,
		Count:   int64(len(components)),
	})
}

// createComponent godoc
//
//	@Summary	Create Dashboard Components
//	@Tags		dashboard
//	@Accept		json
//	@Produce	json
//	@Success	201				{object}	models.Component
//	@Failure	400				{object}	ErrorResponse
//	@Param		id				path		string								true	"Dashboard ID"
//	@Param		Authorization	header		string								true	"Bearer <Access Token>"
//	@Param		input			body		api.createComponent.componentInput	true	"Component Input"
//	@Router		/api/dashboards/{id}/components [post]
func (a *API) createComponent(c echo.Context) error {
	dashboard := c.Get(dashboardContextKey).(*models.Dashboard)

	type componentInput struct {
		Name    string               `json:"name" validate:"required" err:"name is required"`
		Type    models.ComponentType `json:"type" validate:"type" err:"name is required"`
		Configs interface{}          `json:"configs"`
		Query   string               `json:"query" validate:"required" err:"query is required"`
	}

	var in componentInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	component, err := a.db.CreateComponent(dashboard, in.Name, in.Type, in.Configs, in.Query)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, component)
}

// updateComponent godoc
//
//	@Summary	Update Dashboard Component
//	@Tags		dashboard
//	@Accept		json
//	@Produce	json
//	@Success	200				{object}	models.Component
//	@Failure	400				{object}	ErrorResponse
//	@Param		id				path		string								true	"Dashboard ID"
//	@Param		cpid			path		string								true	"Component ID"
//	@Param		input			body		api.updateComponent.componentInput	true	"Component Input"
//	@Param		Authorization	header		string								true	"Bearer <Access Token>"
//	@Router		/api/dashboards/{id}/components/{cpid} [put]
func (a *API) updateComponent(c echo.Context) error {
	component := c.Get(componentContextKey).(*models.Component)

	type componentInput struct {
		Name    string               `json:"name" validate:"required" err:"name is required"`
		Type    models.ComponentType `json:"type" validate:"required" err:"type is required"`
		Configs interface{}          `json:"configs"`
		Query   string               `json:"query" validate:"required" err:"query is required"`
	}

	var in componentInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	component.Name = in.Name
	component.Type = in.Type
	component.Query = in.Query

	if in.Configs != nil {
		cfg, err := json.Marshal(in.Configs)
		if err != nil {
			return err
		}
		component.Configs = cfg
	} else {
		component.Configs = nil
	}

	if err := a.db.UpdateComponent(component); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, component)
}

// deleteComponent godoc
//
//	@Summary	Delete Dashboard Component
//	@Tags		dashboard
//	@Accept		json
//	@Produce	json
//	@Success	204
//	@Param		id				path	string	true	"Dashboard ID"
//	@Param		cpid			path	string	true	"Component ID"
//	@Param		Authorization	header	string	true	"Bearer <Access Token>"
//	@Router		/api/dashboards/{id}/components/{cpid} [delete]
func (a *API) deleteComponent(c echo.Context) error {
	component := c.Get(componentContextKey).(*models.Component)

	if err := a.db.DeleteComponent(component); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// runComponent godoc
//
//	@Summary	Run Dashboard Component
//	@Tags		dashboard
//	@Accept		json
//	@Produce	json
//	@Success	200				{object}	queue.TaskResult
//	@Param		id				path		string	true	"Dashboard ID"
//	@Param		cpid			path		string	true	"Component ID"
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Router		/api/dashboards/{id}/components/{cpid}/run [post]
func (a *API) runComponent(c echo.Context) error {
	component := c.Get(componentContextKey).(*models.Component)

	task := a.q.AddTask(func() (interface{}, error) {
		return a.db.RunQuery(component.Query)
	})

	return c.JSON(http.StatusOK, task)
}
