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
