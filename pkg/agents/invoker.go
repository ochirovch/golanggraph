package agents

import "github.com/ochirovch/golanggraph/pkg/tools"

type Invoker interface {
	Invoke(config Config, messages Messages) Messages
	BindTools(tools []tools.Tool)
	GetTools() []tools.Tool
}
