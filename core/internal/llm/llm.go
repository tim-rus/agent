package llm

import (
	"errors"

	"github.com/openai/openai-go/v3"
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
