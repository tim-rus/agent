package llm

import "github.com/openai/openai-go/v3"

//

func (llm *LLM) MakeParams(model string, msgs []Message) openai.ChatCompletionNewParams {
	return openai.ChatCompletionNewParams{
		Model:    model,
		Messages: messagesToUnion(msgs),
	}
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
