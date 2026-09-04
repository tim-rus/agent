package chat

type Mood string

const (
	MoodNeutral  Mood = "neutral"
	MoodPositive Mood = "positive"
	MoodNegative Mood = "negative"
)

//

type Response struct {
	Text string `json:"text"`
	Mood string `json:"mood"`
}

//

var responseSchema = map[string]string{
	"text": "string",
	"mood": "string",
}
