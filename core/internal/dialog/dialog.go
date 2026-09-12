package dialog

import "core/internal/llm"

type Dialog struct {
	llm *llm.LLM
}

func New(llm *llm.LLM) *Dialog {
	return &Dialog{
		llm: llm,
	}
}
