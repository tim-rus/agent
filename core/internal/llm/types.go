package llm

//

type Role string

const (
	RoleSystem    Role = "system"
	RoleAssistant Role = "assistant"
	RoleUser      Role = "user"
)

//

type Message struct {
	Role    Role
	Content string
}

//

type Usage struct {
	Prompt     int64
	Completion int64
	Total      int64
	Cached     int64
}

type CompletionResponse struct {
	Content string
	Usage   Usage
}
