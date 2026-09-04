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
	usage   Usage
	model   string
}

func NewChat(client openai.Client, model string, systemPromt string) *Chat {
	return &Chat{
		client: client,
		history: []Message{
			{
				Role:    RoleSystem,
				Message: systemPromt,
			},
		},
		usage: Usage{},
		model: model,
	}
}

//

func (chat *Chat) ask(ctx context.Context, q string) (string, error) {
	res, err := chat.request(ctx, q)
	if err != nil {
		slog.Error("request error", "err", err)
		return "", ErrRequest
	}

	chat.usage.Total += uint64(res.Usage.TotalTokens)
	chat.usage.Promts += uint64(res.Usage.PromptTokens)
	chat.usage.Responses += uint64(res.Usage.CompletionTokens)

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

//

func (chat *Chat) request(ctx context.Context, msg string) (*openai.ChatCompletion, error) {
	union := msgsToUnions(chat.history)
	u := append(union, messageToUnion(Message{Role: RoleUser, Message: msg}))

	res, err := chat.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model:    chat.model,
			Messages: u,
		},
		option.WithJSONSet("enable_thinking", false),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to make openai request: %w", err)
	}

	return res, nil
}

// utils

func msgsToUnions(msgs []Message) []openai.ChatCompletionMessageParamUnion {
	list := make([]openai.ChatCompletionMessageParamUnion, len(msgs), len(msgs)+1) // room for new user message
	for i, m := range msgs {
		list[i] = messageToUnion(m)
	}
	return list
}

func messageToUnion(msg Message) openai.ChatCompletionMessageParamUnion {
	switch msg.Role {
	case RoleSystem:
		return openai.SystemMessage(msg.Message)
	case RoleAssistant:
		return openai.AssistantMessage(msg.Message)
	default:
		return openai.UserMessage(msg.Message)
	}
}
