package router

import (
	"github.com/labstack/echo/v4"
)

type Group func(*echo.Group)
