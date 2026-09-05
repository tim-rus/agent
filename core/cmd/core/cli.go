package main

import (
	"bufio"
	"context"
	"core/internal/chat"
	"errors"
	"fmt"
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

//

func cli_loop(ctx context.Context, chatSvc *chat.Chat) error {
	scanner := bufio.NewScanner(os.Stdin)

	printWelcome()

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
			printError(res, err)
			break
		}

		if res.Injection {
			break
		}

		if res.Compact != "" {
			break
		}

		chatSvc.UpdateHistory(input, res.Text)

		printResponse(res)

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

//

func printWelcome() {
	fmt.Println(msgWelcome)
	fmt.Printf("[type '%s' to quit]\n", cmdExit)
}

func printResponse(res *chat.Response) {
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

func printError(res *chat.Response, err error) {
	switch {
	case errors.Is(err, chat.ErrRequest):
		fmt.Println(msgeErrRequest)
	case errors.Is(err, chat.ErrEmpty):
		fmt.Println(msgeErrEmpty)
	}
}
