package agents

type AgentType string

const (
	AgentTypeReact AgentType = "react"
)

type Config struct {
	ThreadID string
}

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
	RoleTool      Role = "tool"
)
