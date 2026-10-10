package main

import (
	"core/internal/gen/pingconnect"
	"log/slog"
	"net"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/labstack/echo/v5"
)

func main() {

	router := setupRouter()

	ping := &Ping{}

	path, handler := pingconnect.NewPingServiceHandler(
		ping,
		connect.WithInterceptors(validate.NewInterceptor()),
	)

	router.Any(path+"*", echo.WrapHandler(handler))

	host := "127.0.0.1"
	port := "9000"

	pingServer := setupHTTPServer(net.JoinHostPort(host, port), router)

	slog.Info("starting core", "host", host, "port", port)

	if err := pingServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("ping server error", "err", err)
		os.Exit(1)
	}
}
