package tools

import (
	"github.com/ochirovch/golanggraph/pkg/agents/message"
	"github.com/ochirovch/golanggraph/pkg/agents/state"
)

type Tool func(parameters map[string]any) (output map[string]any, err error)

type ToolNode interface {
	Call(state state.State) (message.Messages, error)
}
