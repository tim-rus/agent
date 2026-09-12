package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

//

const (
	defaultModel = "qwen/qwen3.7-flash"
)

//

var (
	ErrEmptyResult = errors.New("empty result")
)

//

func main() {
	if err := godotenv.Load("../config/.core.local.env"); err != nil {
		slog.Error("failed to load env file", "err", err)
		os.Exit(1)
	}

	oai := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_KEY")),
		option.WithBaseURL(os.Getenv("OPENAI_BASE_URL")),
	)

	if err := cliLoop(useRequest(oai)); err != nil {
		slog.Error("cli loop failed", "err", err)
	}
}

//

type RequestFunc func(ctx context.Context, q string) (string, error)

func useRequest(client openai.Client) RequestFunc {
	return func(ctx context.Context, q string) (string, error) {
		res, err := client.Chat.Completions.New(
			ctx,
			openai.ChatCompletionNewParams{
				Model: defaultModel,
				Messages: []openai.ChatCompletionMessageParamUnion{
					openai.SystemMessage("будь максимально кратким"),
					openai.UserMessage(q),
				},
			},
			option.WithJSONSet("enable_thinking", false),
		)
		if err != nil {
			slog.Error("REQ ERR", "err", err)
			return "", fmt.Errorf("failed to make openai request: %w", err)
		}

		if len(res.Choices) == 0 {
			return "", ErrEmptyResult
		}

		return res.Choices[0].Message.Content, nil
	}
}
