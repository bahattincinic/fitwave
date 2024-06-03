package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaginatedResponse struct {
	Results interface{} `json:"results"`
	Count   int64       `json:"count"`
}

const (
	maxPageSize = int64(100)
)

// GetPageAndSize retrieves the pagination parameters (page and page size) from the query parameters of the request context.
// If no parameters are provided, it uses the default page size. It also ensures the page and page size values are valid.
// Returns the offset for database queries, the page size, and an error if any of the parameters are invalid.
func (a *API) GetPageAndSize(c echo.Context, defaultPageSize int64) (int64, int64, error) {
	pageSize := defaultPageSize
	if tmp := c.QueryParam("limit"); tmp != "" {
		i, _ := strconv.ParseInt(tmp, 10, 64)
		if i < 1 {
			return 0, 0, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid page_size"))
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
			return 0, 0, echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("invalid page"))
		}
		page = i
	}
	offset := (page - 1) * pageSize

	return offset, pageSize, nil
}
