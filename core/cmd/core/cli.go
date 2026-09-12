package main

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"
)

func cliLoop(request RequestFunc) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		res, err := request(context.Background(), input)
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
