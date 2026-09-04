package chat

import (
	"context"
	"errors"
	"log/slog"

	"github.com/openai/openai-go/v3"
)

//

var (
	ErrRequest = errors.New("request error")
	ErrEmpty   = errors.New("empty result")
)

//

type Role string

const (
	RoleSystem    Role = "system"
	RoleAssistant Role = "assistant"
	RoleUser      Role = "user"
)

type Message struct {
	Role    Role
	Message string
}

type Usage struct {
	Total     uint64
	Promts    uint64
	Responses uint64
}

//

type Chat struct {
	client  openai.Client
	history []Message
	model   string
	Usage   Usage
}

func New(client openai.Client, model string, systemPromt string) *Chat {
	return &Chat{
		client: client,
		history: []Message{
			{
				Role:    RoleSystem,
				Message: systemPromt,
			},
		},
		model: model,
		Usage: Usage{},
	}
}

//

func (chat *Chat) Ask(ctx context.Context, q string) (string, error) {
	res, err := chat.request(ctx, q)
	if err != nil {
		slog.Error("request error", "err", err)
		return "", ErrRequest
	}

	chat.Usage.Total += uint64(res.Usage.TotalTokens)
	chat.Usage.Promts += uint64(res.Usage.PromptTokens)
	chat.Usage.Responses += uint64(res.Usage.CompletionTokens)

	if len(res.Choices) < 1 {
		slog.Warn("request returned empty result", "res", res)
		return "", ErrEmpty
	}

	content := res.Choices[0].Message.Content

	msg := Message{
		Role:    RoleAssistant,
		Message: content,
	}

	chat.history = append(chat.history, msg)

	return content, nil

}
