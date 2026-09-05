package main

import (
	"bufio"
	"context"
	"core/internal/chat"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const (
	msgWelcome = "hey there!"
	msgGoodbye = "bye"

	msgStatusThinking = "thinking..."

	msgeErrRequest = "request error"
	msgeErrEmpty   = "request returned empty result"

	cmdExit = ":q"
)

func cli_loop(ctx context.Context, chatSvc *chat.Chat) error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msgWelcome)
	fmt.Printf("[type '%s' to quit]\n", cmdExit)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if strings.TrimSpace(input) == cmdExit {
			fmt.Printf("tokens spent: %d\n", chatSvc.Usage.Total)
			fmt.Println(msgGoodbye)
			break
		}

		fmt.Println(msgStatusThinking)

		res, err := chatSvc.Ask(ctx, input)
		if err != nil {
			switch {
			case errors.Is(err, chat.ErrRequest):
				fmt.Println(msgeErrRequest)
			case errors.Is(err, chat.ErrEmpty):
				fmt.Println(msgeErrEmpty)
			}
			break
		}

		slog.Info("RES", "res", res)

		fmt.Println(res.Text)
		fmt.Println("mood:", res.Mood)
		if res.Compact != "" {
			fmt.Println("summury")
			fmt.Println(res.Compact)
		}
		if res.Injection {
			fmt.Println("WARNING: injection!")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
