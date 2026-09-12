package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../config/.core.local.env"); err != nil {
		slog.Error("failed to load env file", "err", err)
		os.Exit(1)
	}

	if err := cliLoop(); err != nil {
		slog.Error("cli loop failed", "err", err)
	}
}
