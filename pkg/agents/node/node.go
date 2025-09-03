package node

import "github.com/ochirovch/golanggraph/pkg/agents"

type Node struct {
	Name     string
	LLM      agents.Invoker
	Function NodeFunc
}

type Edge struct {
	From string
	To   string
}

type NodeFunc func(llm agents.Invoker, messages agents.Messages) (
	retMessages agents.Messages, data map[string]any)
