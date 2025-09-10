package message

import (
	"fmt"

	"github.com/ochirovch/golanggraph/pkg/agents"
)

type ToolCall struct {
	Name string
	Args map[string]any
	ID   string
	Type string
}

type Messages []Message
type Message struct {
	Role      agents.Role
	Content   string
	ToolCalls []ToolCall
}

func (m Messages) Print() string {
	var result string
	for _, msg := range m {
		result += fmt.Sprintf("[%s]: %s\n", msg.Role, msg.Content)
	}
	return result
}
