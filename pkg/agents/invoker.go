package agents

import "github.com/ochirovch/golanggraph/pkg/tools"

type Invoker interface {
	Invoke(config Config, messages Messages, tools []tools.Tool) Messages
}
