package main

import (
	"core/internal/gen/pingconnect"
	"core/internal/platform/arguments"
	"core/internal/platform/config"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/labstack/echo/v5"
)

//

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

//

func main() {
	args := arguments.Parse()

	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: args.LogLevel})))

	cfg := Config{}
	if err := config.Read(&cfg, os.Getenv("CONFIG_FILE"), os.Getenv("ENV_FILE")); err != nil {
		slog.Error("failed to read configuration", "error", err)
		os.Exit(1)
	}

	router := setupRouter()

	ping := &Ping{}

	path, handler := pingconnect.NewPingServiceHandler(
		ping,
		connect.WithInterceptors(validate.NewInterceptor()),
	)

	router.Any(path+"*", echo.WrapHandler(handler))

	pingServer := setupHTTPServer(net.JoinHostPort(cfg.Host, fmt.Sprint(cfg.Port)), router)

	slog.Info("starting core", "host", cfg.Host, "port", cfg.Port)

	if err := pingServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("ping server error", "err", err)
		os.Exit(1)
	}
}
