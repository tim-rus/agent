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
	params openai.ChatCompletionNewParams
}

func New(client openai.Client, model string) *LLM {
	return &LLM{
		client: client,
		params: openai.ChatCompletionNewParams{
			Model: model,
		},
	}
}

//

func (llm *LLM) Request(ctx context.Context, msgs []Message, thinking bool) (CompletionResponse, error) {
	p := llm.params
	p.Messages = messagesToUnion(msgs)

	res, err := llm.client.Chat.Completions.New(
		ctx,
		p,
		option.WithJSONSet("enable_thinking", thinking),
	)
	if err != nil {
		return CompletionResponse{}, fmt.Errorf("failed to make openai request: %w", err)
	}

	if len(res.Choices) == 0 {
		return CompletionResponse{}, ErrEmptyResult
	}

	return CompletionResponse{
		Content: res.Choices[0].Message.Content,
		Usage: Usage{
			Prompt:     res.Usage.PromptTokens,
			Completion: res.Usage.CompletionTokens,
			Total:      res.Usage.TotalTokens,
			Cached:     res.Usage.PromptTokensDetails.CachedTokens,
		},
	}, nil
}

//

func messagesToUnion(msgs []Message) []openai.ChatCompletionMessageParamUnion {
	union := make([]openai.ChatCompletionMessageParamUnion, len(msgs))
	for i, m := range msgs {
		switch m.Role {
		case RoleSystem:
			union[i] = openai.SystemMessage(m.Content)
		case RoleAssistant:
			union[i] = openai.AssistantMessage(m.Content)
		default:
			union[i] = openai.UserMessage(m.Content)
		}
	}
	return union
}
