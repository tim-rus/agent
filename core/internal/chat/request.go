package chat

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/shared"
)

func (chat *Chat) request(ctx context.Context, msg string) (*openai.ChatCompletion, error) {
	union := msgsToUnions(chat.history)
	u := append(union, messageToUnion(Message{Role: RoleUser, Message: fmt.Sprintf("<user_input>\n%s\n</user_input>", msg)}))

	res, err := chat.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model:    chat.model,
			Messages: u,
			ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
				OfJSONSchema: &shared.ResponseFormatJSONSchemaParam{
					JSONSchema: shared.ResponseFormatJSONSchemaJSONSchemaParam{
						Schema: chat.responseSchema,
						Strict: openai.Bool(true),
					},
				},
			},
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
