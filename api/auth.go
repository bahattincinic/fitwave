package api

import (
	"net/http"
	"time"

	"github.com/bahattincinic/fitwave/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	pkgerrors "github.com/pkg/errors"
)

// decodeToken decodes and validates a JWT token string.
// Returns true if the token is valid and not expired, otherwise returns false.
func (a *API) decodeToken(s string) bool {
	token, err := jwt.ParseWithClaims(s, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.cfg.API.SecretKey), nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims.VerifyExpiresAt(time.Now().Unix(), true)
	}

	return false
}

// createToken creates a JWT token with an expiry time based on the API configuration.
// Returns the signed token as a string and an error if the token creation or signing fails.
func (a *API) createToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(a.cfg.API.TokenExpiryHour) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(a.cfg.API.SecretKey))

	if err != nil {
		return "", err
	}
	return s, nil
}

// login godoc
//
//	@Summary	Login
//	@Tags		auth
//	@Accept		json
//	@Param		input	body		api.login.loginInput	true	"Login Input"
//	@Success	201		{object}	map[string]string
//	@Failure	400		{object}	ErrorResponse
//	@Router		/api/auth/token [post]
func (a *API) login(c echo.Context) error {
	type loginInput struct {
		Username string `json:"username" validate:"required" err:"username is required"`
		Password string `json:"password" validate:"required" err:"password is required"`
	}

	var in loginInput
	if err := a.bindAndValidate(c, &in); err != nil {
		return err
	}

	cfg, err := a.db.GetCurrentConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if !cfg.SetupCompleted() {
		return echo.NewHTTPError(http.StatusBadRequest,
			pkgerrors.New("setup is not completed yet"))
	}

	if cfg.LoginType != models.ProtectedLoginType {
		return echo.NewHTTPError(http.StatusBadRequest,
			pkgerrors.New("Only protected login needs access token"))
	}

	if !cfg.CheckLogin(in.Username, in.Password, a.cfg.API.SecretKey) {
		return echo.NewHTTPError(http.StatusBadRequest,
			pkgerrors.New("username or password is invalid"))
	}

	accessToken, err := a.createToken()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token": accessToken,
	})
}
