package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Created return with `Created` status code and created entity schema json.
func Created(c echo.Context, schema interface{}) error {
	return c.JSON(http.StatusCreated, schema)
}
