package api

import (
	"net/http"

	"github.com/bahattincinic/fitwave/models"
	"github.com/labstack/echo/v4"
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

// updateConfig godoc
//
//	@Summary	Upsert Application Config
//	@Tags		config
//	@Accept		json
//	@Param		config			body		api.updateConfig.updateInput	true	"Config Input"
//	@Param		Authorization	header		string							true	"Bearer <Access Token>"
//	@Success	200				{object}	models.Config
//	@Failure	400				{object}	ErrorResponse
//	@Router		/api/config [put]
func (a *API) updateConfig(c echo.Context) error {
	type updateInput struct {
		ClientId     int    `json:"client_id" validate:"min=1" err:"client_id is required"`
		ClientSecret string `json:"client_secret" validate:"required" err:"client_secret is required"`
	}

	var in updateInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	cfg.ClientId = in.ClientId
	cfg.ClientSecret = in.ClientSecret

	if _, err = a.db.UpsertConfig(*cfg); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, cfg)
}

// checkSetupCompleted godoc
//
//	@Summary	Check Application Setup completed
//	@Tags		config
//	@Accept		json
//	@Success	200	{object}	map[string]interface{}
//	@Router		/api/config/setup [get]
func (a *API) checkSetupCompleted(c echo.Context) error {
	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"completed":  cfg.SetupCompleted(),
		"login_type": cfg.LoginType,
	})
}

// completeSetup godoc
//
//	@Summary	Complete Application Setup
//	@Tags		config
//	@Accept		json
//	@Param		input	body		api.completeSetup.setupInput	true	"Setup Input"
//	@Success	201		{object}	map[string]bool
//	@Failure	400		{object}	ErrorResponse
//	@Router		/api/config/setup [post]
func (a *API) completeSetup(c echo.Context) error {
	type setupInput struct {
		ClientId      int              `json:"client_id" validate:"min=1" err:"client_id is required"`
		ClientSecret  string           `json:"client_secret" validate:"required" err:"client_secret is required"`
		LoginUsername string           `json:"login_username"`
		LoginPassword string           `json:"login_password"`
		LoginType     models.LoginType `json:"login_type" validate:"oneof=anonymous protected" err:"login_type is required"`
	}

	var in setupInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	if in.LoginType == models.ProtectedLoginType &&
		(in.LoginUsername == "" || in.LoginPassword == "") {
		return echo.NewHTTPError(http.StatusBadRequest,
			"username and password are required for protected login")
	}

	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cfg.SetupCompleted() {
		return echo.NewHTTPError(http.StatusBadRequest, "setup is already completed")
	}

	cfg.ClientId = in.ClientId
	cfg.ClientSecret = in.ClientSecret
	cfg.LoginType = in.LoginType

	if in.LoginType == models.ProtectedLoginType {
		cfg.LoginUsername = in.LoginUsername
		cfg.SetPassword(in.LoginPassword, a.cfg.API.SecretKey)
	}

	cfg, err = a.db.UpsertConfig(*cfg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, map[string]bool{
		"status": cfg.SetupCompleted(),
	})
}
