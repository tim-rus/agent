package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func setupRouter() *echo.Echo {
	r := echo.NewWithConfig(echo.Config{}) // TODO: config

	r.Pre(middleware.RemoveTrailingSlash())
	r.Use(middleware.Recover())
	r.Use(middleware.RequestID())

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		UnsafeAllowOriginFunc: func(c *echo.Context, origin string) (string, bool, error) {
			return origin, true, nil // TODO: change for production
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"Connect-Protocol-Version",
		},
		ExposeHeaders: []string{
			"Connect-Context-Accept-Encoding",
			"Connect-Content-Encoding",
		},
		AllowCredentials: true,
	}))

	return r
}
