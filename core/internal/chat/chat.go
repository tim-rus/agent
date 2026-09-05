package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/openai/openai-go/v3"
)

//

var (
	ErrRequest   = errors.New("request error")
	ErrEmpty     = errors.New("empty result")
	ErrUnmarshal = errors.New("decoding error")
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
	client openai.Client

	model          string
	history        []Message
	responseSchema map[string]any

	Usage Usage
}

func New(client openai.Client, model string, systemPromt string) (*Chat, error) {
	schema, err := generateSchema[Response]()
	if err != nil {
		return nil, fmt.Errorf("failed to generate schema: %w", err)
	}

	return &Chat{
		client: client,

		model:          model,
		responseSchema: schema,

		history: []Message{
			{
				Role:    RoleSystem,
				Message: systemPromt,
			},
		},

		Usage: Usage{},
	}, nil
}

//

func (chat *Chat) Ask(ctx context.Context, q string) (*Response, error) {
	res, err := chat.request(ctx, q)
	if err != nil {
		slog.Error("request error", "err", err)
		return nil, ErrRequest
	}

	chat.updateUsage(res.Usage)

	if len(res.Choices) < 1 {
		slog.Warn("request returned empty result", "res", res)
		return nil, ErrEmpty
	}

	content := res.Choices[0].Message.Content

	userMsg := Message{
		Role:    RoleUser,
		Message: q,
	}

	assistantMsg := Message{
		Role:    RoleAssistant,
		Message: content,
	}

	chat.history = append(chat.history, userMsg)
	chat.history = append(chat.history, assistantMsg)

	response := Response{}
	if err := json.Unmarshal([]byte(content), &response); err != nil {
		slog.Error("unmarshalling error", "err", err)
		return nil, ErrEmpty
	}

	return &response, nil
}

func (chat *Chat) updateUsage(usage openai.CompletionUsage) {
	chat.Usage.Total += uint64(usage.TotalTokens)
	chat.Usage.Promts += uint64(usage.PromptTokens)
	chat.Usage.Responses += uint64(usage.CompletionTokens)
}

//
