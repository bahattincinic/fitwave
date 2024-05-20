package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// getAccessToken godoc
//
//	@Summary	Get Access Token from Auth Code
//	@Tags		auth
//	@Accept		json
//	@Param		auth	body		api.getAccessToken.tokenRequestInput	true	"Access Token Input"
//	@Success	200		{object}	strava.AuthorizationResponse
//	@Router		/auth/token [post]
func (a *API) getAccessToken(c echo.Context) error {
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

// getAccessToken godoc
//
//	@Summary	Get Access Token from Auth Code
//	@Tags		auth
//	@Accept		json
//	@Param		callback_url	query		string	true	"Callback URL"
//	@Success	200				{object}	map[string]string
//	@Router		/auth/authorization-url [get]
func (a *API) getAuthorizationURL(c echo.Context) error {
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
