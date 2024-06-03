package api

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/bahattincinic/fitwave/models"
	"github.com/bahattincinic/fitwave/strava"
	"github.com/labstack/echo/v4"
)

// getStravaAccessToken godoc
//
//	@Summary	Get Strava Access Token from Auth Code
//	@Tags		strava
//	@Accept		json
//	@Param		auth	body		api.getStravaAccessToken.tokenRequestInput	true	"Access Token Input"
//	@Success	200		{object}	strava.AuthorizationResponse
//	@Router		/api/strava/token [post]
func (a *API) getStravaAccessToken(c echo.Context) error {
	type tokenRequestInput struct {
		Code string `json:"code"`
	}

	var in tokenRequestInput
	if err := c.Bind(&in); err != nil {
		return err
	}

	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resp, err := a.st.GetAccessToken(cfg, in.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// getStravaAccessToken godoc
//
//	@Summary	Get Authorization URL for Strava Login
//	@Tags		strava
//	@Accept		json
//	@Param		callback_url	query		string	true	"Callback URL"
//	@Success	200				{object}	map[string]string
//	@Router		/api/strava/authorization-url [get]
func (a *API) getStravaAuthorizationURL(c echo.Context) error {
	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	callbackURL := c.QueryParam("callback_url")
	if callbackURL == "" {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Errorf("callback_url should be provided"))
	}

	url := a.st.GetAuthorizationURL(cfg, callbackURL)

	return c.JSON(http.StatusOK, map[string]string{
		"authorization_url": url,
	})
}

// getStravaMe godoc
//
//	@Summary	Get Current Strava User Details
//	@Tags		strava
//	@Accept		json
//	@Param		X-Strava-Authorization	header		string	true	"Strava Access Token"
//	@Success	200						{object}	strava.User
//	@Router		/api/strava/me [get]
func (a *API) getStravaMe(c echo.Context) error {
	user := c.Get(userContextKey).(*strava.User)
	return c.JSON(http.StatusOK, user)
}

// syncData godoc
//
//	@Summary	Sync Strava data
//	@Tags		strava
//	@Accept		json
//	@Param		X-Strava-Authorization	header		string	true	"Strava Access Token"
//	@Success	200						{object}	queue.TaskResult
//	@Router		/api/strava/sync [post]
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

// exportActivityGPS godoc
//
//	@Summary	Export Activity GPX
//	@Tags		strava
//	@Accept		json
//	@Param		id						path	string	true	"Activity ID"
//	@Param		X-Strava-Authorization	header	string	true	"Strava Access Token"
//	@Success	200
//	@Router		/api/strava/activities/{id}/gpx [get]
func (a *API) exportActivityGPS(c echo.Context) error {
	user := c.Get(userContextKey).(*strava.User)
	act := c.Get(activityContextKey).(*models.Activity)

	gpx, err := a.st.ExportGPX(user, act.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if gpx == "" {
		return echo.NewHTTPError(http.StatusNotFound, "GPX Not Found")
	}

	c.Response().Header().Set(echo.HeaderContentDisposition,
		fmt.Sprintf("attachment; filename=\"activity_%d.gpx\"", act.Id))
	c.Response().Header().Set(echo.HeaderContentType,
		"application/gpx+xml")
	return c.Blob(http.StatusOK, "application/gpx+xml", []byte(xml.Header+gpx))
}

// getActivityLaps godoc
//
//	@Summary	Get Activity Laps
//	@Tags		strava
//	@Accept		json
//	@Param		id						path		string	true	"Activity ID"
//	@Param		X-Strava-Authorization	header		string	true	"Strava Access Token"
//	@Success	200						{object}	PaginatedResponse{Results=[]strava.LapEffortSummary, count=int}
//	@Router		/api/strava/activities/{id}/laps [get]
func (a *API) getActivityLaps(c echo.Context) error {
	user := c.Get(userContextKey).(*strava.User)
	act := c.Get(activityContextKey).(*models.Activity)

	laps, err := a.st.GetActivityLaps(user, act.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, PaginatedResponse{
		Results: laps,
		Count:   int64(len(laps)),
	})
}
