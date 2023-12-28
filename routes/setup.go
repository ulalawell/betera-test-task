package routes

import "github.com/labstack/echo/v4"

func SetupRoutes(apodHandler *Apod, e *echo.Echo) {
	e.GET("api/pictures", apodHandler.getApods, RateLimiter())
	e.GET("api/picture/:date", apodHandler.getApodsByDate, RateLimiter())
}
