package api

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type PaginatedResponse struct {
	Results interface{} `json:"results"`
	Count   int64       `json:"count"`
}

type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

const (
	maxPageSize = int64(100)
	errTag      = "err"
)

// GetPageAndSize retrieves the pagination parameters (page and page size) from the query parameters of the request context.
// If no parameters are provided, it uses the default page size. It also ensures the page and page size values are valid.
// Returns the offset for database queries, the page size, and an error if any of the parameters are invalid.
func (a *API) GetPageAndSize(c echo.Context, defaultPageSize int64) (int64, int64, error) {
	pageSize := defaultPageSize
	if tmp := c.QueryParam("limit"); tmp != "" {
		i, _ := strconv.ParseInt(tmp, 10, 64)
		if i < 1 {
			return 0, 0, echo.NewHTTPError(http.StatusBadRequest, "invalid page_size")
		}

		if i < maxPageSize {
			pageSize = i
		} else {
			pageSize = maxPageSize
		}
	}

	page := int64(1)
	if tmp := c.QueryParam("page"); tmp != "" {
		i, _ := strconv.ParseInt(tmp, 10, 64)
		if i < 1 {
			return 0, 0, echo.NewHTTPError(http.StatusBadRequest, "invalid page")
		}
		page = i
	}
	offset := (page - 1) * pageSize

	return offset, pageSize, nil
}

// bindAndValidate binds the data from the request to the provided target interface
// and validates it using the validator.
func (a *API) bindAndValidate(c echo.Context, target interface{}) error {
	if err := c.Bind(target); err != nil {
		return err
	}

	if err := a.val.Struct(target); err != nil {
		errorMessages := make(map[string]string)
		for _, vErr := range err.(validator.ValidationErrors) {
			fieldName := vErr.Field()
			s := reflect.ValueOf(target).Elem().Interface()
			field, _ := reflect.TypeOf(s).FieldByName(fieldName)
			errorMessages[field.Tag.Get("json")] = field.Tag.Get(errTag)
		}
		errorResponse := ErrorResponse{Errors: errorMessages}
		return echo.NewHTTPError(http.StatusBadRequest, errorResponse)
	}
	return nil
}
