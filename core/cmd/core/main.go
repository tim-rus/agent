package main

import "log/slog"

func main() {
	if err := cliLoop(); err != nil {
		slog.Error("cli loop failed", "err", err)
	}
}
