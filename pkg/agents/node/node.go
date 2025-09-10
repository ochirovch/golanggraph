package node

import (
	"github.com/ochirovch/golanggraph/pkg/agents/invoker"
	"github.com/ochirovch/golanggraph/pkg/agents/message"
	"github.com/ochirovch/golanggraph/pkg/agents/state"
)

type Node struct {
	Name     string
	LLM      invoker.Invoker
	Function NodeFunc
	Tool     NodeTool
}

type NodeFunc func(llm invoker.Invoker, messages message.Messages) (
	retMessages message.Messages, data map[string]any)

type NodeTool interface {
	Call(state state.State) (message.Messages, error)
}
