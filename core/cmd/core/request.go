package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

//

var (
	ErrEmptyResult = errors.New("empty result")
)

//

type RequestFunc func(ctx context.Context, q string) (string, error)

//

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
