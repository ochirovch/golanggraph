package invoker

import (
	"github.com/ochirovch/golanggraph/pkg/agents"
	"github.com/ochirovch/golanggraph/pkg/agents/message"
	"github.com/ochirovch/golanggraph/pkg/agents/tools"
)

type Invoker interface {
	Invoke(config agents.Config, messages message.Messages) message.Messages
	BindTools(tools []tools.Tool)
	GetTools() []tools.Tool
}
