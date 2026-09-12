package main

import (
	"bufio"
	"context"
	"core/internal/dialog"
	"fmt"
	"log/slog"
	"os"
)

func cliLoop(d *dialog.Dialog) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		res, err := d.Ask(context.Background(), input)
		if err != nil {
			slog.Error("request error", "err", err)
			println("Request error")
			continue
		}

		println()
		println(res)
		println()
		println()

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
