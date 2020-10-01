package error

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewBadRequestError creates a new error with http.StatusBadRequest
func NewBadRequestError(message ...interface{}) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusBadRequest, message...)
}

// NewUnauthorizedError creates a new error with http.StatusUnauthorized
func NewUnauthorizedError(message ...interface{}) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnauthorized, message...)
}

// NewValidationError creates a new error with http.StatusUnprocessableEntity
func NewValidationError(message ...interface{}) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, message...)
}
