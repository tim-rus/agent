package main

import (
	"net/http"

	connectcors "connectrpc.com/cors"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func setupRouter() *echo.Echo {
	r := echo.NewWithConfig(echo.Config{}) // TODO: config

	r.Pre(middleware.RemoveTrailingSlash())
	r.Use(middleware.Recover())
	r.Use(middleware.RequestID())

	r.OPTIONS("/*", func(c *echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		UnsafeAllowOriginFunc: func(c *echo.Context, origin string) (string, bool, error) {
			return origin, true, nil // TODO: change for production
		},
		AllowMethods:     connectcors.AllowedMethods(), // GET, POST, OPTIONS
		AllowHeaders:     append(connectcors.AllowedHeaders(), "Origin", "Content-Type", "Accept", "Authorization"),
		ExposeHeaders:    connectcors.ExposedHeaders(),
		AllowCredentials: true,
	}))

	return r
}
