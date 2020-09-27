package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	kitmw "github.com/Huangkai1008/micro-kit/pkg/middleware"
	"github.com/Huangkai1008/micro-kit/pkg/validator"
)

type Group func(*echo.Group)

// NewRouter returns a new Echo router.
func NewRouter(logger *zap.Logger, group Group, validator *validator.CustomValidator) (*echo.Echo, error) {
	e := echo.New()
	e.Logger.SetHeader("${time_rfc3339} ${level} ${prefix} ${short_file} ${line}")
	e.Use(middleware.Recover())
	e.Use(kitmw.LoggerMiddleware(logger))
	e.Validator = validator
	apiGroup := e.Group("/api")
	v1Group := apiGroup.Group("/v1")
	{
		group(v1Group)
	}
	return e, nil
}
