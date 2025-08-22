package agents

import "fmt"

type Messages []Message
type Message struct {
	Role    Role
	Content string
}

func (m Messages) Print() string {
	var result string
	for _, msg := range m {
		result += fmt.Sprintf("[%s]: %s\n", msg.Role, msg.Content)
	}
	return result
}
