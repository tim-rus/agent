package dialog

import "core/internal/llm"

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
