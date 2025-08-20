package agents

import "golanggraph/pkg/tools"

type Invoker interface {
	Invoke(config Config, messages []Message, tools []tools.Tool) string
}
