package api

import (
	"net/http"

	"github.com/bahattincinic/fitwave/models"
	"github.com/labstack/echo/v4"
	pkgerrors "github.com/pkg/errors"
)

// getConfig godoc
//
//	@Summary	Get Application Config
//	@Tags		config
//	@Accept		json
//	@Param		Authorization	header		string	true	"Bearer <Access Token>"
//	@Success	200				{object}	models.Config
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
//	@Tags		config
//	@Accept		json
//	@Param		config			body		models.Config	true	"Config Input"
//	@Param		Authorization	header		string			true	"Bearer <Access Token>"
//	@Success	200				{object}	models.Config
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

// checkSetupCompleted godoc
//
//	@Summary	Check Application Setup completed
//	@Tags		config
//	@Accept		json
//	@Success	200	{object}	map[string]bool
//	@Router		/api/config/setup [get]
func (a *API) checkSetupCompleted(c echo.Context) error {
	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]bool{
		"status": cfg.SetupCompleted(),
	})
}

// completeSetup godoc
//
//	@Summary	Complete Application Setup
//	@Tags		config
//	@Accept		json
//	@Param		input	body		api.completeSetup.setupInput	true	"Setup Input"
//	@Success	201		{object}	map[string]bool
//	@Router		/api/config/setup [post]
func (a *API) completeSetup(c echo.Context) error {
	type setupInput struct {
		ClientId      int              `json:"client_id"`
		ClientSecret  string           `json:"client_secret"`
		LoginUsername string           `json:"login_username"`
		LoginPassword string           `json:"login_password"`
		LoginType     models.LoginType `json:"login_type"`
	}

	var in setupInput
	if err := c.Bind(&in); err != nil {
		return err
	}

	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cfg.SetupCompleted() {
		return echo.NewHTTPError(http.StatusBadRequest,
			pkgerrors.New("setup is already completed"))
	}

	cfg.ClientId = in.ClientId
	cfg.ClientSecret = in.ClientSecret
	cfg.LoginType = in.LoginType
	cfg.LoginUsername = in.LoginUsername
	cfg.LoginPassword = in.LoginPassword

	cfg, err = a.db.UpsertConfig(*cfg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]bool{
		"status": cfg.SetupCompleted(),
	})
}
