package llm

import (
	"context"
	"errors"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

//

var (
	ErrEmptyResult = errors.New("empty result")
)

//

type LLM struct {
	client openai.Client
}

func New(client openai.Client) *LLM {
	return &LLM{
		client: client,
	}
}

//

func (llm *LLM) Request(ctx context.Context, params openai.ChatCompletionNewParams, thinking bool) (openai.ChatCompletionChoice, error) {
	res, err := llm.client.Chat.Completions.New(
		ctx,
		params,
		option.WithJSONSet("enable_thinking", thinking),
	)
	if err != nil {
		return openai.ChatCompletionChoice{}, fmt.Errorf("failed to make openai request: %w", err)
	}

	if len(res.Choices) == 0 {
		return openai.ChatCompletionChoice{}, ErrEmptyResult
	}

	return res.Choices[0], nil
}
