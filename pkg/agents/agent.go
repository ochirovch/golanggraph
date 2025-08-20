package agents

import "golanggraph/pkg/tools"

type AgentType string

const (
	AgentTypeReact AgentType = "react"
)

type Agent interface {
	Invoke(config Config, messages []Message, tools []tools.Tool) string
}

type Config struct {
	ThreadID string
}

type Messages []Message
type Message struct {
	Role    Role
	Content string
}

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
	RoleTool      Role = "tool"
)
