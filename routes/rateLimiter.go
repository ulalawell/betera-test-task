package routes

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
	"time"
)

func RateLimiter() echo.MiddlewareFunc {
	limit := rate.NewLimiter(rate.Every(time.Second), 10)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limit.Allow() {
				return echo.ErrTooManyRequests
			}
			return next(c)
		}
	}
}
