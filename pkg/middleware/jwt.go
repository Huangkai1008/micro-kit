package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/Huangkai1008/micro-kit/pkg/auth"
	e "github.com/Huangkai1008/micro-kit/pkg/error"
)

// AuthMiddleware is the middleware function for authorization.
func AuthMiddleware(j auth.Auth) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return e.ErrAccountEmptyAuthHeader
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 {
				return e.ErrAccountInvalidAuthHeader
			}

			tokenString := parts[1]
			claims, err := j.ParseToken(tokenString)
			if err != nil {
				return e.ErrInvalidToken
			}

			if claims.HasExpired() {
				return e.ErrTokenExpired
			}

			c.Set("identity", claims.GetIdentity())
			return next(c)
		}
	}
}
