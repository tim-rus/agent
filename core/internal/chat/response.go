package chat

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

//

type Mood string

const (
	MoodNeutral  Mood = "neutral"
	MoodPositive Mood = "positive"
	MoodNegative Mood = "negative"
)

//

type Response struct {
	Text    string `json:"text"`
	Mood    string `json:"mood"`
	Compact string `json:"compact,omitempty"`
}

//

// copy(!) from openai-go package readme
func generateSchema[T any]() (map[string]any, error) {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	data, err := json.Marshal(schema)
	if err != nil {
		return nil, fmt.Errorf("marshal JSON schema: %w", err)
	}
	var rawSchema map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawSchema); err != nil {
		return nil, fmt.Errorf("decode JSON schema: %w", err)
	}
	result := make(map[string]any, len(rawSchema))
	for key, value := range rawSchema {
		result[key] = value
	}
	return result, nil
}
