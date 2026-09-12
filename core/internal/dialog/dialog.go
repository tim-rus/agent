package dialog

import (
	"context"
	"core/internal/llm"
	"fmt"
)

//

type Dialog struct {
	llm     *llm.LLM
	history []llm.Message
	usage   llm.Usage
}

func New(l *llm.LLM, systemPrompt string) *Dialog {
	return &Dialog{
		llm: l,
		history: []llm.Message{
			{
				Role:    llm.RoleSystem,
				Content: systemPrompt,
			},
		},
		usage: llm.Usage{},
	}
}

//

func (d *Dialog) Ask(ctx context.Context, q string) (string, error) {
	msg := llm.Message{Role: llm.RoleUser, Content: q}

	// copy history in case of request error
	msgs := make([]llm.Message, len(d.history)+1)
	copy(msgs, d.history)
	msgs[len(d.history)] = msg

	res, err := d.llm.Request(ctx, msgs, false)
	if err != nil {
		return "", fmt.Errorf("failed to request llm: %w", err)
	}

	d.history = append(d.history, msg, llm.Message{Role: llm.RoleAssistant, Content: res.Content})

	return res.Content, nil
}
