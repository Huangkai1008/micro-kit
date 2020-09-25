package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			statusCode := res.Status

			fields := []zapcore.Field{
				zap.String("time", time.Now().Format(time.RFC3339Nano)),
				zap.String("remote_ip", c.RealIP()),
				zap.String("host", req.Host),
				zap.String("method", req.Method),
				zap.String("request_uri", req.RequestURI),
				zap.String("user_agent", req.UserAgent()),
				zap.Int("status", res.Status),
				zap.String("latency", stop.Sub(start).String()),
			}
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
				fields = append(fields, zap.String("request_id", id))
			}

			switch {
			case statusCode >= 400 && statusCode <= 499:
				{
					logger.With(zap.Error(err)).Warn("[WARN]", fields...)
				}
			case statusCode >= 500:
				{
					logger.With(zap.Error(err)).Error("[ERROR]", fields...)
				}
			default:
				logger.Info("[INFO]", fields...)
			}
			return nil
		}
	}
}
