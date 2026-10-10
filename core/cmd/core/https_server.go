package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func setupHTTPServer(addr string, router *echo.Echo) *http.Server {
	protocols := &http.Protocols{}
	protocols.SetUnencryptedHTTP2(true)
	protocols.SetHTTP1(true)

	httpServer := &http.Server{
		Addr:      addr,
		Handler:   router,
		Protocols: protocols,
	} // TODO: config (timeouts)

	return httpServer
}
