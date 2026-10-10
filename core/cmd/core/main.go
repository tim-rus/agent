package main

import (
	"context"
	"core/internal/gen/pingconnect"
	"core/internal/platform/arguments"
	"core/internal/platform/config"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/SladkyCitron/slogcolor"
	"github.com/labstack/echo/v5"
)

//

const defaultShutdownTimeoutSeconds = 15

//

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	ShutdownTimeoutSeconds int `yaml:"shutdown_timeout_seconds"`
}

//

func main() {
	args := arguments.Parse()

	opts := slogcolor.DefaultOptions
	opts.Level = args.LogLevel

	slog.SetDefault(slog.New(
		slogcolor.NewHandler(os.Stderr, opts),
	))

	cfg := Config{}
	if err := config.Read(&cfg, os.Getenv("CONFIG_FILE"), os.Getenv("ENV_FILE")); err != nil {
		slog.Error("failed to read configuration", "error", err)
		os.Exit(1)
	}

	if cfg.ShutdownTimeoutSeconds == 0 {
		cfg.ShutdownTimeoutSeconds = defaultShutdownTimeoutSeconds
	}

	ctx, sysCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer sysCancel()

	if err := run(ctx, cfg); err != nil {
		slog.Error("app failed", "error", err)
		os.Exit(1)
	}

	slog.Info("app exited")
}

//

func run(ctx context.Context, cfg Config) error {

	router := setupRouter()

	ping := &Ping{}

	path, handler := pingconnect.NewPingServiceHandler(
		ping,
		connect.WithInterceptors(validate.NewInterceptor()),
	)

	router.Any(path+"*", echo.WrapHandler(handler))

	pingServer := setupHTTPServer(net.JoinHostPort(cfg.Host, fmt.Sprint(cfg.Port)), router)

	// run

	apiErrCh := make(chan error, 1)

	go func() {
		slog.Info("api server started", "at", pingServer.Addr)
		if err := pingServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			apiErrCh <- err
		}
	}()

	// shutdown

	var unexpectedErr error

	select {
	case unexpectedErr = <-apiErrCh:
		slog.Error("api server failed", "error", unexpectedErr)
	case <-ctx.Done():
		slog.Info("interrupt signal received")
	}

	slog.Info("shutting down")

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.ShutdownTimeoutSeconds)*time.Second,
	)
	defer shutdownCancel()

	if err := pingServer.Shutdown(shutdownCtx); err != nil {
		slog.Error("failed to gracefully shutdown api server", "error", err)
	}

	return unexpectedErr
}
